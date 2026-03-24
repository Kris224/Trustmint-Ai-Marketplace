const { ethers } = require("ethers");

async function main() {
  const provider = new ethers.JsonRpcProvider("https://polygon-amoy.drpc.org/");
  const address = "0xc21b3a355Bec5E6a7E02fC993DE32917769eB172";
  
  console.log(`Checking address: ${address} on Polygon Amoy...`);
  
  try {
    const code = await provider.getCode(address);
    if (code === "0x") {
      console.log("❌ ERROR: NO CONTRACT FOUND AT THIS ADDRESS ON AMOY (Bytecode is 0x)!");
    } else {
      console.log(`✅ CONTRACT EXISTS! Bytecode length: ${code.length}`);
      
      const abi = ["function totalSupply() view returns (uint256)"];
      const contract = new ethers.Contract(address, abi, provider);
      
      try {
        const supply = await contract.totalSupply();
        console.log(`✅ totalSupply() call succeeded! Value: ${supply.toString()}`);
      } catch(e) {
        console.log(`❌ totalSupply() call FAILED: ${e.message}`);
      }
    }
  } catch (e) {
    console.log("RPC Error:", e.message);
  }
}

main();
