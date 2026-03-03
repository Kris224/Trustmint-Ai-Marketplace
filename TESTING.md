# Testing Guide - Wallet-Bound CLI Download

## Prerequisites

### 1. Install MetaMask Browser Extension
https://metamask.io/download/

### 2. Import Hardhat Test Account to MetaMask
- Open MetaMask → Click account icon → "Import Account"
- Paste private key:
```
0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
```

### 3. Add Hardhat Network to MetaMask
- MetaMask → Networks → Add Network → Add Manually
- **Network Name**: Hardhat Local
- **RPC URL**: http://127.0.0.1:8545
- **Chain ID**: 31337
- **Currency Symbol**: ETH
- Click "Save" and switch to this network

---

## Start All Services

### Terminal 1: Hardhat Node (Already Running)
```bash
cd ~/Desktop/Trustmint-Ai-Marketplace/blockchain
npx hardhat node
```

### Terminal 2: Frontend (Already Running)
```bash
cd ~/Desktop/Trustmint-Ai-Marketplace/frontend
npm run dev
```

### Terminal 3: Backend (Already Running)
```bash
cd ~/Desktop/Trustmint-Ai-Marketplace/trustmint-backend
source venv/bin/activate
python app.py
```

---

## Test the Wallet-Bound CLI Download

### Step 1: Open Developer Page
Open browser: http://localhost:5173/developer

### Step 2: Connect Wallet
1. Click "🦊 Connect MetaMask" button
2. In MetaMask popup: Click "Next" → "Connect"
3. Verify wallet address shows: `0xf39f...266c`

### Step 3: Download CLI
1. Click "📥 Download CLI" button
2. In MetaMask popup: Click "Sign"
3. CLI downloads as `trustmint` file

### Step 4: Verify Download
```bash
cd ~/Downloads
ls -lh trustmint
chmod +x trustmint
./trustmint --help
```

### Step 5: Check Generated Config
```bash
ls -lh /tmp/cli-downloads/
cat /tmp/cli-downloads/trustmint-0xf39Fd6.config.json
```

Expected config:
```json
{
  "walletAddress": "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
  "network": "hardhat",
  "rpcUrl": "http://127.0.0.1:8545",
  "contracts": {
    "nft": "0x5FbDB2315678afecb367f032d93F642f64180aa3",
    "marketplace": "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
  }
}
```

---

## Test Error Cases

### Test 1: Download Without Wallet
1. Refresh page (disconnects wallet)
2. Try to download
3. Expected: Error toast "Please connect your wallet first!"

### Test 2: Reject Signature
1. Connect wallet
2. Click Download CLI
3. Click "Reject" in MetaMask
4. Expected: Error toast "Signature rejected"

### Test 3: Backend Signature Verification
```bash
curl -X POST http://localhost:5001/api/cli/download \
  -H "Content-Type: application/json" \
  -d '{
    "walletAddress": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "signature": "0xfakesignature",
    "message": "test"
  }'
```
Expected: `{"error": "Invalid signature"}`

---

## Verify Services Running

Check all services are up:
```bash
# Check Hardhat
curl http://127.0.0.1:8545 -X POST -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'

# Check Frontend
curl http://localhost:5173

# Check Backend
curl http://localhost:5001/api/cli/download -X OPTIONS
```

---

## Troubleshooting

### Backend not starting?
```bash
cd ~/Desktop/Trustmint-Ai-Marketplace/trustmint-backend
source venv/bin/activate
pip install flask-cors web3 eth-account
python app.py
```

### CLI binary missing?
```bash
cd ~/Desktop/Trustmint-Ai-Marketplace/trustmint-cli
go build -o trustmint
```

### CORS errors?
Check backend terminal for errors, ensure flask-cors is installed

---

## Success Checklist

- [ ] All 3 services running (Hardhat, Frontend, Backend)
- [ ] MetaMask connected on Hardhat network
- [ ] Wallet shows correct address
- [ ] CLI downloads successfully
- [ ] Config file has correct wallet address
- [ ] `./trustmint --help` command works
