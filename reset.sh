#!/bin/bash
# reset.sh — Run this every time you restart npx hardhat node
# Usage: ./reset.sh (from the project root)

set -e

PROJECT_ROOT="$(cd "$(dirname "$0")" && pwd)"
BLOCKCHAIN_DIR="$PROJECT_ROOT/blockchain"
BACKEND_ENV="$PROJECT_ROOT/trustmint-backend/.env"
FRONTEND_ENV="$PROJECT_ROOT/frontend/.env"

echo "🚀 Deploying contracts to local Hardhat node..."
cd "$BLOCKCHAIN_DIR"
npx hardhat run scripts/deploy.cjs --network localhost

# Read new addresses from localhost.json
NFT_ADDR=$(node -e "const f=require('./deployments/localhost.json'); console.log(f.contracts.TrustmintNFT)")
MKT_ADDR=$(node -e "const f=require('./deployments/localhost.json'); console.log(f.contracts.TrustmintMarketplace)")

echo ""
echo "📝 New contract addresses:"
echo "   TrustmintNFT:         $NFT_ADDR"
echo "   TrustmintMarketplace: $MKT_ADDR"

# Update backend .env
echo ""
echo "⚙️  Updating trustmint-backend/.env..."
sed -i "s|^NFT_CONTRACT=.*|NFT_CONTRACT=$NFT_ADDR|" "$BACKEND_ENV"
sed -i "s|^MARKETPLACE_CONTRACT=.*|MARKETPLACE_CONTRACT=$MKT_ADDR|" "$BACKEND_ENV"

# Update frontend .env
echo "⚙️  Updating frontend/.env..."
sed -i "s|^VITE_NFT_ADDRESS=.*|VITE_NFT_ADDRESS=$NFT_ADDR|" "$FRONTEND_ENV"
sed -i "s|^VITE_MARKETPLACE_ADDRESS=.*|VITE_MARKETPLACE_ADDRESS=$MKT_ADDR|" "$FRONTEND_ENV"

echo ""
echo "✅ Done! Contract addresses updated in both .env files."
echo ""
echo "👉 Next steps:"
echo "   1. Restart: python app.py   (in trustmint-backend/)"
echo "   2. Restart: npm run dev     (in frontend/)"
echo "   3. MetaMask → Settings → Advanced → Clear activity and nonce data"
echo ""
