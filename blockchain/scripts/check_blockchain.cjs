const hre = require("hardhat");
const fs = require("fs");
const path = require("path");

async function main() {
    console.log("🔍 Checking Blockchain Data...\n");

    // Load deployment info
    const networkName = hre.network.name;
    const deploymentFile = path.join(__dirname, `../deployments/${networkName}.json`);
    if (!fs.existsSync(deploymentFile)) {
        console.error(`❌ No deployment found. Please run: npx hardhat run scripts/deploy.cjs --network ${networkName}`);
        return;
    }

    const deployment = JSON.parse(fs.readFileSync(deploymentFile, "utf8"));
    const nftAddress = deployment.contracts.TrustmintNFT;
    const marketplaceAddress = deployment.contracts.TrustmintMarketplace;

    console.log("📋 Deployment Info:");
    console.log("  Network:", deployment.network);
    console.log("  Deployer:", deployment.deployer);
    console.log("  TrustmintNFT:", nftAddress);
    console.log("  Marketplace:", marketplaceAddress);
    console.log("  Deployed at block:", deployment.blockNumber);
    console.log("");

    // Get contract instances
    const TrustmintNFT = await hre.ethers.getContractFactory("TrustmintNFT");
    const nft = TrustmintNFT.attach(nftAddress);

    const TrustmintMarketplace = await hre.ethers.getContractFactory("TrustmintMarketplace");
    const marketplace = TrustmintMarketplace.attach(marketplaceAddress);

    // Check NFT data
    console.log("━".repeat(60));
    console.log("🎨 TrustmintNFT Data");
    console.log("━".repeat(60));

    const totalSupply = await nft.totalSupply();
    console.log(`Total Models Minted: ${totalSupply}\n`);

    if (totalSupply > 0) {
        for (let i = 1; i <= totalSupply; i++) {
            try {
                const owner = await nft.ownerOf(i);
                const metadata = await nft.getModelMetadata(i);
                const tokenURI = await nft.tokenURI(i);

                console.log(`📦 Token ID: ${i}`);
                console.log(`  Owner: ${owner}`);
                console.log(`  Creator: ${metadata.creator}`);
                console.log(`  Model Hash: ${metadata.modelHash}`);
                console.log(`  Dataset Hash: ${metadata.datasetHash}`);
                console.log(`  Merkle Root: ${metadata.merkleRoot}`);
                console.log(`  IPFS CID: ${metadata.ipfsCid}`);
                console.log(`  Verified: ${metadata.verified}`);
                console.log(`  Minted: ${new Date(Number(metadata.timestamp) * 1000).toISOString()}`);
                console.log(`  Token URI: ${tokenURI}`);
                console.log("");
            } catch (error) {
                console.log(`  ⚠️  Token ${i} error: ${error.message}`);
            }
        }
    } else {
        console.log("  ℹ️  No models have been minted yet.");
    }

    // Check Marketplace data
    console.log("\n" + "━".repeat(60));
    console.log("🛒 TrustmintMarketplace Data");
    console.log("━".repeat(60));

    const totalListings = await marketplace.totalListings();
    const platformFee = await marketplace.platformFee();
    console.log(`Platform Fee: ${Number(platformFee) / 100}%`);
    console.log(`Total Listings Created: ${totalListings}\n`);

    if (totalListings > 0) {
        for (let i = 1; i <= totalListings; i++) {
            try {
                const listing = await marketplace.getListing(i);

                console.log(`📋 Listing ID: ${i}`);
                console.log(`  Token ID: ${listing.tokenId}`);
                console.log(`  NFT Contract: ${listing.nftContract}`);
                console.log(`  Seller: ${listing.seller}`);
                console.log(`  Price: ${hre.ethers.formatEther(listing.price)} ETH`);
                console.log(`  Status: ${listing.active ? "🟢 Active" : "🔴 Inactive"}`);
                console.log("");
            } catch (error) {
                console.log(`  ⚠️  Listing ${i} error: ${error.message}`);
            }
        }
    } else {
        console.log("  ℹ️  No listings have been created yet.");
    }

    console.log("━".repeat(60));
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
