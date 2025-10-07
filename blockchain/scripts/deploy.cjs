
const hre = require("hardhat");

async function main() {
  const Registry = await hre.ethers.getContractFactory("ModelRegistry");
  const registry = await Registry.deploy();
  await (registry.waitForDeployment ? registry.waitForDeployment() : registry.deployed());
  console.log("ModelRegistry deployed to:", registry.target ? registry.target : registry.address);
}

main().catch((e) => { console.error(e); process.exit(1); });
