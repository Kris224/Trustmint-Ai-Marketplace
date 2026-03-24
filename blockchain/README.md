# Trustmint Blockchain - Setup Guide

## Overview

The Trustmint blockchain infrastructure uses **Hardhat** for smart contract development and testing, with **ERC721 NFTs** to represent AI models and a **marketplace contract** for trading.

## Smart Contracts

### 1. TrustmintNFT.sol
- **Type**: ERC721 NFT with metadata storage
- **Purpose**: Each AI model is minted as a unique NFT
- **Features**:
  - On-chain storage of model hash, dataset hash, and merkle root
  - IPFS CID storage for large files
  - Creator royalties (5% via ERC2981)
  - Verification functions
  - Duplicate prevention

### 2. TrustmintMarketplace.sol
- **Purpose**: Decentralized marketplace for buying/selling model NFTs
- **Features**:
  - Listing creation and cancellation
  - Purchase with native cryptocurrency
  - Platform fees (2.5%)
  - Royalty enforcement
  - Safe escrow mechanism

## Setup Instructions

### Prerequisites
```bash
cd blockchain
npm install
```

### Environment Configuration

Create a `.env` file (use `.env.example` as template):
```bash
# For local development - no keys needed
# For testnet deployment:
PRIVATE_KEY=your_wallet_private_key
POLYGON_AMOY_RPC=https://rpc-amoy.polygon.technology
POLYGONSCAN_API_KEY=your_api_key
PINATA_JWT=your_pinata_jwt
```

### Compile Contracts
```bash
npx hardhat compile
```

### Run Tests
```bash
npx hardhat test
```

All tests should pass:
- ✅ 45+ test cases
- ✅ Minting, verification, marketplace, royalties
- ✅ Gas estimation included

### Deploy Locally

**Terminal 1 - Start Local Network:**
```bash
npx hardhat node
```

**Terminal 2 - Deploy Contracts:**
```bash
npx hardhat run scripts/deploy.cjs --network hardhat
```

Deployed addresses will be saved to `deployments/hardhat.json`.

### Deploy to Testnet

```bash
# Polygon Amoy Testnet
npx hardhat run scripts/deploy.cjs --network polygon_amoy

# Verify on PolygonScan
npx hardhat verify --network polygon_amoy <NFT_CONTRACT_ADDRESS>
npx hardhat verify --network polygon_amoy <MARKETPLACE_ADDRESS>
```

## Contract Interactions

### Minting a Model NFT

```javascript
const modelHash = "0x1234..."; // SHA-256 of model file
const datasetHash = "0xabcd..."; // SHA-256 of dataset
const merkleRoot = "0x5678..."; // Merkle root of training proof
const ipfsCid = "QmXxx..."; // IPFS CID
const metadataUri = "ipfs://QmMetadata..."; // Metadata URI

const tx = await nftContract.mintModel(
  modelHash,
  datasetHash,
  merkleRoot,
  ipfsCid,
  metadataUri
);

const receipt = await tx.wait();
console.log("NFT Minted! Token ID:", receipt.events[0].args.tokenId);
```

### Listing on Marketplace

```javascript
// Approve marketplace
await nftContract.setApprovalForAll(marketplaceAddress, true);

// List for sale
const price = ethers.parseEther("1.0"); // 1 ETH/MATIC
await marketplaceContract.listModel(nftAddress, tokenId, price);
```

### Purchasing a Model

```javascript
const listingId = 1;
const price = ethers.parseEther("1.0");

await marketplaceContract.purchaseModel(listingId, { value: price });
console.log("Model purchased!");
```

## CLI Integration

The CLI uses Go to interact with smart contracts:

### Configuration
Edit `trustmint-cli/blockchain-config.json`:
```json
{
  "network": "hardhat",
  "rpcUrl": "http://127.0.0.1:8545",
  "contracts": {
    "TrustmintNFT": "0x5FbDB2315678afecb367f032d93F642f64180aa3",
    "TrustmintMarketplace": "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
  }
}
```

### Blockchain Packages

- **`internal/blockchain/client.go`**: Ethereum client wrapper
- **`internal/blockchain/ipfs.go`**: Pinata/IPFS integration
- **`internal/blockchain/nft.go`**: NFT contract interactions (to be completed)

## IPFS Configuration

### Using Pinata

1. Sign up at https://pinata.cloud
2. Get API credentials (JWT recommended)
3. Set environment variables:
   ```bash
   PINATA_JWT=your_jwt_here
   ```

### Upload Example

```go
import "trustmint-cli/internal/blockchain"

config := blockchain.PinataConfig{
    JWT: os.Getenv("PINATA_JWT"),
}

// Upload file
cid, err := blockchain.UploadToPinata("model.pth", config)

// Upload JSON metadata
metadata := map[string]interface{}{
    "name": "My AI Model",
    "modelHash": "0x1234...",
}
cid, err := blockchain.UploadJSONToPinata(metadata, "metadata.json", config)
```

## Gas Costs (Estimated)

Based on hardhat tests with optimizer enabled:

| Operation | Gas Used | Approx Cost (Polygon) |
|-----------|----------|----------------------|
| Mint NFT | ~372k | ~$0.01 |
| List for Sale | ~216k | <$0.01 |
| Purchase | ~108k | <$0.01 |
| Deploy NFT | ~3.7M | ~$0.10 |
| Deploy Marketplace | ~2.3M | ~$0.06 |

*Costs are approximate and vary with gas prices*

## Next Steps

1. **Complete CLI Integration**:
   - Add `trustmint publish` blockchain functionality
   - Implement MetaMask signature verification
   - Auto-upload to IPFS on publish

2. **Generate Contract Bindings**:
   ```bash
   # Install abigen
   go install github.com/ethereum/go-ethereum/cmd/abigen@latest
   
   # Generate Go bindings
   abigen --abi artifacts/contracts/TrustmintNFT.sol/TrustmintNFT.json \
          --pkg contracts --type TrustmintNFT --out internal/contracts/nft.go
   ```

3. **Backend Integration**:
   - Add endpoints for CLI download with wallet binding
   - Add NFT verification endpoints
   - Add IPFS metadata fetching

4. **Frontend** (Deferred):
   - MetaMask connection
   - NFT display
   - Marketplace UI

## Troubleshooting

### "Module not found" errors
Make sure you're in the `blockchain/` directory and ran `npm install`.

### Test failures
- Ensure you're using Node.js v16+ (v18+ recommended)
- Run `npx hardhat clean` then recompile

### Deployment issues
- Check `.env` file has correct values
- Ensure wallet has sufficient funds
- Verify RPC URL is accessible

## Resources

- [Hardhat Documentation](https://hardhat.org/docs)
- [OpenZeppelin Contracts](https://docs.openzeppelin.com/contracts)
- [Pinata IPFS](https://docs.pinata.cloud)
- [Polygon Amoy Faucet](https://faucet.polygon.technology)
