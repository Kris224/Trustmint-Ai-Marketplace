const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("TrustmintMarketplace", function () {
    let nft, marketplace;
    let owner, seller, buyer, creator;
    let tokenId;

    beforeEach(async function () {
        [owner, seller, buyer, creator] = await ethers.getSigners();

        // Deploy contracts
        const TrustmintNFT = await ethers.getContractFactory("TrustmintNFT");
        nft = await TrustmintNFT.deploy();
        await nft.waitForDeployment();

        const TrustmintMarketplace = await ethers.getContractFactory("TrustmintMarketplace");
        marketplace = await TrustmintMarketplace.deploy();
        await marketplace.waitForDeployment();

        await nft.setMarketplace(await marketplace.getAddress());

        // Mint a test NFT
        const tx = await nft.connect(seller).mintModel(
            "0x1234567890abcdef",
            "0xabcdef1234567890",
            ethers.keccak256(ethers.toUtf8Bytes("test")),
            "QmTest123",
            "ipfs://metadata"
        );
        await tx.wait();
        tokenId = 1;

        // Approve marketplace
        await nft.connect(seller).setApprovalForAll(await marketplace.getAddress(), true);
    });

    describe("Listing", function () {
        it("Should create a listing", async function () {
            const price = ethers.parseEther("1.0");
            const listingTx = await marketplace.connect(seller).listModel(
                await nft.getAddress(),
                tokenId,
                price
            );

            const receipt = await listingTx.wait();
            const event = receipt.logs.find(log => {
                try {
                    return marketplace.interface.parseLog(log).name === "ListingCreated";
                } catch (e) {
                    return false;
                }
            });

            expect(event).to.not.be.undefined;

            const listing = await marketplace.getListing(1);
            expect(listing.seller).to.equal(seller.address);
            expect(listing.price).to.equal(price);
            expect(listing.active).to.equal(true);
        });

        it("Should not allow listing without approval", async function () {
            await nft.connect(seller).setApprovalForAll(await marketplace.getAddress(), false);
            const price = ethers.parseEther("1.0");

            await expect(
                marketplace.connect(seller).listModel(
                    await nft.getAddress(),
                    tokenId,
                    price
                )
            ).to.be.revertedWith("Marketplace not approved");
        });

        it("Should not allow listing if not owner", async function () {
            const price = ethers.parseEther("1.0");

            await expect(
                marketplace.connect(buyer).listModel(
                    await nft.getAddress(),
                    tokenId,
                    price
                )
            ).to.be.revertedWith("Not the owner of this NFT");
        });

        it("Should not allow zero price", async function () {
            await expect(
                marketplace.connect(seller).listModel(
                    await nft.getAddress(),
                    tokenId,
                    0
                )
            ).to.be.revertedWith("Price must be greater than zero");
        });

        it("Should not allow duplicate listings", async function () {
            const price = ethers.parseEther("1.0");
            await marketplace.connect(seller).listModel(
                await nft.getAddress(),
                tokenId,
                price
            );

            await expect(
                marketplace.connect(seller).listModel(
                    await nft.getAddress(),
                    tokenId,
                    price
                )
            ).to.be.revertedWith("Token already listed");
        });
    });

    describe("Canceling Listing", function () {
        beforeEach(async function () {
            const price = ethers.parseEther("1.0");
            await marketplace.connect(seller).listModel(
                await nft.getAddress(),
                tokenId,
                price
            );
        });

        it("Should allow seller to cancel listing", async function () {
            await marketplace.connect(seller).cancelListing(1);
            const listing = await marketplace.getListing(1);
            expect(listing.active).to.equal(false);
        });

        it("Should not allow non-seller to cancel", async function () {
            await expect(
                marketplace.connect(buyer).cancelListing(1)
            ).to.be.revertedWith("Not the seller");
        });

        it("Should not allow canceling inactive listing", async function () {
            await marketplace.connect(seller).cancelListing(1);
            await expect(
                marketplace.connect(seller).cancelListing(1)
            ).to.be.revertedWith("Listing not active");
        });
    });

    describe("Purchasing", function () {
        const listingPrice = ethers.parseEther("1.0");

        beforeEach(async function () {
            await marketplace.connect(seller).listModel(
                await nft.getAddress(),
                tokenId,
                listingPrice
            );
        });

        it("Should allow purchase with exact price", async function () {
            const sellerBalanceBefore = await ethers.provider.getBalance(seller.address);

            await marketplace.connect(buyer).purchaseModel(1, {
                value: listingPrice
            });

            // Check NFT transferred
            expect(await nft.ownerOf(tokenId)).to.equal(buyer.address);

            // Check listing deactivated
            const listing = await marketplace.getListing(1);
            expect(listing.active).to.equal(false);

            // Check seller received payment (minus fees)
            const sellerBalanceAfter = await ethers.provider.getBalance(seller.address);
            const platformFee = (listingPrice * BigInt(250)) / BigInt(10000); // 2.5%
            const royalty = (listingPrice * BigInt(500)) / BigInt(10000); // 5%
            const expectedProceeds = listingPrice - platformFee - royalty;

            expect(sellerBalanceAfter - sellerBalanceBefore).to.be.closeTo(
                expectedProceeds,
                ethers.parseEther("0.01") // Allow small variance for gas
            );
        });

        it("Should refund excess payment", async function () {
            const overpayment = ethers.parseEther("2.0");
            const buyerBalanceBefore = await ethers.provider.getBalance(buyer.address);

            const tx = await marketplace.connect(buyer).purchaseModel(1, {
                value: overpayment
            });
            const receipt = await tx.wait();
            const gasUsed = receipt.gasUsed * receipt.gasPrice;

            const buyerBalanceAfter = await ethers.provider.getBalance(buyer.address);

            // Buyer should have spent approximately the listing price + gas
            const totalSpent = buyerBalanceBefore - buyerBalanceAfter;
            expect(totalSpent).to.be.closeTo(
                listingPrice + gasUsed,
                ethers.parseEther("0.01")
            );
        });

        it("Should not allow purchase with insufficient payment", async function () {
            await expect(
                marketplace.connect(buyer).purchaseModel(1, {
                    value: ethers.parseEther("0.5")
                })
            ).to.be.revertedWith("Insufficient payment");
        });

        it("Should not allow seller to buy own NFT", async function () {
            await expect(
                marketplace.connect(seller).purchaseModel(1, {
                    value: listingPrice
                })
            ).to.be.revertedWith("Seller cannot buy own NFT");
        });

        it("Should not allow purchase of inactive listing", async function () {
            await marketplace.connect(seller).cancelListing(1);

            await expect(
                marketplace.connect(buyer).purchaseModel(1, {
                    value: listingPrice
                })
            ).to.be.revertedWith("Listing not active");
        });
    });

    describe("Platform Fee", function () {
        it("Should allow owner to update platform fee", async function () {
            const newFee = 500; // 5%
            await marketplace.connect(owner).setPlatformFee(newFee);
            expect(await marketplace.platformFee()).to.equal(newFee);
        });

        it("Should not allow fee greater than 10%", async function () {
            await expect(
                marketplace.connect(owner).setPlatformFee(1001)
            ).to.be.revertedWith("Fee cannot exceed 10%");
        });

        it("Should not allow non-owner to update fee", async function () {
            await expect(
                marketplace.connect(seller).setPlatformFee(500)
            ).to.be.reverted;
        });
    });

    describe("Fee Withdrawal", function () {
        it("Should allow owner to withdraw accumulated fees", async function () {
            // Create and complete a purchase to generate fees
            const price = ethers.parseEther("1.0");
            await marketplace.connect(seller).listModel(
                await nft.getAddress(),
                tokenId,
                price
            );

            await marketplace.connect(buyer).purchaseModel(1, { value: price });

            const ownerBalanceBefore = await ethers.provider.getBalance(owner.address);
            const contractBalance = await ethers.provider.getBalance(await marketplace.getAddress());

            if (contractBalance > 0) {
                const tx = await marketplace.connect(owner).withdrawFees();
                const receipt = await tx.wait();
                const gasUsed = receipt.gasUsed * receipt.gasPrice;

                const ownerBalanceAfter = await ethers.provider.getBalance(owner.address);
                expect(ownerBalanceAfter).to.be.closeTo(
                    ownerBalanceBefore + contractBalance - gasUsed,
                    ethers.parseEther("0.01")
                );
            }
        });

        it("Should not allow non-owner to withdraw fees", async function () {
            await expect(
                marketplace.connect(seller).withdrawFees()
            ).to.be.reverted;
        });
    });

    describe("Utility Functions", function () {
        it("Should return total listings", async function () {
            expect(await marketplace.totalListings()).to.equal(0);

            await marketplace.connect(seller).listModel(
                await nft.getAddress(),
                tokenId,
                ethers.parseEther("1.0")
            );

            expect(await marketplace.totalListings()).to.equal(1);
        });

        it("Should check if token is listed", async function () {
            expect(await marketplace.isListed(await nft.getAddress(), tokenId)).to.equal(false);

            await marketplace.connect(seller).listModel(
                await nft.getAddress(),
                tokenId,
                ethers.parseEther("1.0")
            );

            expect(await marketplace.isListed(await nft.getAddress(), tokenId)).to.equal(true);
        });
    });
});
