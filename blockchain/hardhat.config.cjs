require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();
require("@nomicfoundation/hardhat-verify");

// Default values for local development
const PRIVATE_KEY = process.env.PRIVATE_KEY || "0x0000000000000000000000000000000000000000000000000000000000000000";
const POLYGON_AMOY_RPC = process.env.POLYGON_AMOY_RPC || "";
const POLYGONSCAN_API_KEY = process.env.POLYGONSCAN_API_KEY || "";

module.exports = {
  solidity: {
    compilers: [
      {
        version: "0.8.20",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200
          }
        }
      },
      { version: "0.8.28" },
    ],
  },
  networks: {
    hardhat: {
      chainId: 31337,
    },
    // Only include polygon_amoy if RPC URL is provided
    ...(POLYGON_AMOY_RPC && {
      polygon_amoy: {
        url: POLYGON_AMOY_RPC,
        accounts: [PRIVATE_KEY],
        chainId: 80002
      }
    })
  },
  etherscan: {
    apiKey: POLYGONSCAN_API_KEY
  },
};

