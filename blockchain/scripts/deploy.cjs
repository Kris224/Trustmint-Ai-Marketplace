const hre = require("hardhat");
const fs = require("fs");
const path = require("path");

async function main() {
    console.log("🚀 Starting Trustmint contract deployment...\n");

    // Get deployer account
    const [deployer] = await hre.ethers.getSigners();
    console.log("📝 Deploying contracts with account:", deployer.address);

    const balance = await hre.ethers.provider.getBalance(deployer.address);
    console.log("💰 Account balance:", hre.ethers.formatEther(balance), "ETH\n");

    // Deploy TrustmintNFT
    console.log("📦 Deploying TrustmintNFT contract...");
    const TrustmintNFT = await hre.ethers.getContractFactory("TrustmintNFT");
    const nft = await TrustmintNFT.deploy();
    await nft.waitForDeployment();
    const nftAddress = await nft.getAddress();
    console.log("✅ TrustmintNFT deployed to:", nftAddress);

    // Deploy TrustmintMarketplace
    console.log("\n📦 Deploying TrustmintMarketplace contract...");
    const TrustmintMarketplace = await hre.ethers.getContractFactory("TrustmintMarketplace");
    const marketplace = await TrustmintMarketplace.deploy();
    await marketplace.waitForDeployment();
    const marketplaceAddress = await marketplace.getAddress();
    console.log("✅ TrustmintMarketplace deployed to:", marketplaceAddress);

    // Configure NFT contract with marketplace address
    console.log("\n⚙️  Configuring contracts...");
    const setMarketplaceTx = await nft.setMarketplace(marketplaceAddress);
    await setMarketplaceTx.wait();
    console.log("✅ Marketplace address set in NFT contract");

    // Save deployment addresses
    const deploymentInfo = {
        network: hre.network.name,
        deployer: deployer.address,
        contracts: {
            TrustmintNFT: nftAddress,
            TrustmintMarketplace: marketplaceAddress
        },
        timestamp: new Date().toISOString(),
        blockNumber: await hre.ethers.provider.getBlockNumber()
    };

    const deploymentsDir = path.join(__dirname, "../deployments");
    if (!fs.existsSync(deploymentsDir)) {
        fs.mkdirSync(deploymentsDir, { recursive: true });
    }

    const deploymentFile = path.join(
        deploymentsDir,
        `${hre.network.name}.json`
    );
    fs.writeFileSync(deploymentFile, JSON.stringify(deploymentInfo, null, 2));

    console.log("\n📄 Deployment info saved to:", deploymentFile);

    console.log("\n" + "=".repeat(60));
    console.log("🎉 Deployment Summary");
    console.log("=".repeat(60));
    console.log("Network:           ", hre.network.name);
    console.log("TrustmintNFT:      ", nftAddress);
    console.log("Marketplace:       ", marketplaceAddress);
    console.log("=".repeat(60));

    // Instructions for next steps
    console.log("\n📋 Next Steps:");
    console.log("1. Update CLI configuration with contract addresses");
    console.log("2. Run tests: npx hardhat test");
    console.log("3. Verify contracts (if on testnet):");
    console.log(`   npx hardhat verify --network ${hre.network.name} ${nftAddress}`);
    console.log(`   npx hardhat verify --network ${hre.network.name} ${marketplaceAddress}`);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
