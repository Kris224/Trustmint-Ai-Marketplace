const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("TrustmintNFT", function () {
    let nft, marketplace;
    let owner, creator, buyer;
    let modelHash, datasetHash, merkleRoot, ipfsCid, metadataUri;

    beforeEach(async function () {
        // Get signers
        [owner, creator, buyer] = await ethers.getSigners();

        // Deploy NFT contract
        const TrustmintNFT = await ethers.getContractFactory("TrustmintNFT");
        nft = await TrustmintNFT.deploy();
        await nft.waitForDeployment();

        // Deploy Marketplace contract
        const TrustmintMarketplace = await ethers.getContractFactory("TrustmintMarketplace");
        marketplace = await TrustmintMarketplace.deploy();
        await marketplace.waitForDeployment();

        // Set marketplace in NFT contract
        await nft.setMarketplace(await marketplace.getAddress());

        // Sample test data
        modelHash = "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef";
        datasetHash = "0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890";
        merkleRoot = ethers.keccak256(ethers.toUtf8Bytes("merkle_root_test"));
        ipfsCid = "QmTzQ1234567890abcdefghijklmnopqrstuvwxyz";
        metadataUri = "ipfs://QmMetadata1234567890abcdefghijklmnopqr";
    });

    describe("Deployment", function () {
        it("Should set the correct name and symbol", async function () {
            expect(await nft.name()).to.equal("TrustmintAI");
            expect(await nft.symbol()).to.equal("TMAI");
        });

        it("Should set the correct owner", async function () {
            expect(await nft.owner()).to.equal(owner.address);
        });

        it("Should set marketplace address", async function () {
            expect(await nft.marketplace()).to.equal(await marketplace.getAddress());
        });
    });

    describe("Minting", function () {
        it("Should mint a new model NFT", async function () {
            const tx = await nft.connect(creator).mintModel(
                modelHash,
                datasetHash,
                merkleRoot,
                ipfsCid,
                metadataUri
            );

            const receipt = await tx.wait();
            const event = receipt.logs.find(log => {
                try {
                    return nft.interface.parseLog(log).name === "ModelMinted";
                } catch (e) {
                    return false;
                }
            });

            expect(event).to.not.be.undefined;

            // Check total supply
            expect(await nft.totalSupply()).to.equal(1);

            // Check owner
            expect(await nft.ownerOf(1)).to.equal(creator.address);
        });

        it("Should store correct metadata", async function () {
            await nft.connect(creator).mintModel(
                modelHash,
                datasetHash,
                merkleRoot,
                ipfsCid,
                metadataUri
            );

            const metadata = await nft.getModelMetadata(1);
            expect(metadata.modelHash).to.equal(modelHash);
            expect(metadata.datasetHash).to.equal(datasetHash);
            expect(metadata.merkleRoot).to.equal(merkleRoot);
            expect(metadata.ipfsCid).to.equal(ipfsCid);
            expect(metadata.creator).to.equal(creator.address);
            expect(metadata.verified).to.equal(true);
        });

        it("Should not allow duplicate model hashes", async function () {
            await nft.connect(creator).mintModel(
                modelHash,
                datasetHash,
                merkleRoot,
                ipfsCid,
                metadataUri
            );

            await expect(
                nft.connect(buyer).mintModel(
                    modelHash,
                    "different_dataset_hash",
                    merkleRoot,
                    "different_cid",
                    "different_uri"
                )
            ).to.be.revertedWith("Model already registered");
        });

        it("Should reject empty model hash", async function () {
            await expect(
                nft.connect(creator).mintModel(
                    "",
                    datasetHash,
                    merkleRoot,
                    ipfsCid,
                    metadataUri
                )
            ).to.be.revertedWith("Model hash cannot be empty");
        });

        it("Should reject empty IPFS CID", async function () {
            await expect(
                nft.connect(creator).mintModel(
                    modelHash,
                    datasetHash,
                    merkleRoot,
                    "",
                    metadataUri
                )
            ).to.be.revertedWith("IPFS CID cannot be empty");
        });
    });

    describe("Verification", function () {
        beforeEach(async function () {
            await nft.connect(creator).mintModel(
                modelHash,
                datasetHash,
                merkleRoot,
                ipfsCid,
                metadataUri
            );
        });

        it("Should verify correct model hash", async function () {
            expect(await nft.verifyModel(1, modelHash)).to.equal(true);
        });

        it("Should reject incorrect model hash", async function () {
            expect(await nft.verifyModel(1, "wrong_hash")).to.equal(false);
        });

        it("Should allow owner to update verification status", async function () {
            await nft.connect(owner).setVerificationStatus(1, false);
            const metadata = await nft.getModelMetadata(1);
            expect(metadata.verified).to.equal(false);
        });

        it("Should not allow non-owner to update verification", async function () {
            await expect(
                nft.connect(creator).setVerificationStatus(1, false)
            ).to.be.reverted;
        });
    });

    describe("Token URI and Metadata", function () {
        it("Should return correct token URI", async function () {
            await nft.connect(creator).mintModel(
                modelHash,
                datasetHash,
                merkleRoot,
                ipfsCid,
                metadataUri
            );

            expect(await nft.tokenURI(1)).to.equal(metadataUri);
        });

        it("Should revert for non-existent token", async function () {
            await expect(nft.tokenURI(999)).to.be.reverted;
        });
    });

    describe("Royalties (ERC2981)", function () {
        it("Should support ERC2981 interface", async function () {
            // ERC2981 interface ID
            const ERC2981InterfaceId = "0x2a55205a";
            expect(await nft.supportsInterface(ERC2981InterfaceId)).to.equal(true);
        });

        it("Should return correct royalty info", async function () {
            await nft.connect(creator).mintModel(
                modelHash,
                datasetHash,
                merkleRoot,
                ipfsCid,
                metadataUri
            );

            const salePrice = ethers.parseEther("1.0");
            const royaltyInfo = await nft.royaltyInfo(1, salePrice);

            expect(royaltyInfo[0]).to.equal(creator.address); // Creator receives royalty
            expect(royaltyInfo[1]).to.equal(salePrice * BigInt(500) / BigInt(10000)); // 5%
        });
    });

    describe("Access Control", function () {
        it("Should allow owner to set marketplace", async function () {
            const newMarketplace = buyer.address;
            await nft.connect(owner).setMarketplace(newMarketplace);
            expect(await nft.marketplace()).to.equal(newMarketplace);
        });

        it("Should not allow non-owner to set marketplace", async function () {
            await expect(
                nft.connect(creator).setMarketplace(buyer.address)
            ).to.be.reverted;
        });

        it("Should reject zero address for marketplace", async function () {
            await expect(
                nft.connect(owner).setMarketplace(ethers.ZeroAddress)
            ).to.be.revertedWith("Invalid marketplace address");
        });
    });
});
