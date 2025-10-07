const hre = require("hardhat");

async function main() {
  const argv = process.argv.slice(2);
  if (argv.length < 5) {
    console.log("Usage: node register_model.cjs <registry_address> <modelHash> <datasetHash> <ipfsCid> <merkleRootHex>");
    process.exit(1);
  }
  const [registryAddress, modelHash, datasetHash, ipfsCid, merkleRootHex] = argv;

  const Registry = await hre.ethers.getContractFactory("ModelRegistry");
  const registry = await Registry.attach(registryAddress);

  // merkleRootHex must be 0x... and 32 bytes
  const merkleBytes = hre.ethers.utils.arrayify(merkleRootHex);
  const tx = await registry.registerModel(modelHash, datasetHash, ipfsCid, merkleBytes);
  console.log("Submitted tx, waiting for confirmation...");
  const receipt = await tx.wait();
  console.log("Transaction mined:", receipt.transactionHash);
}

main().catch((e) => { console.error(e); process.exit(1); });
