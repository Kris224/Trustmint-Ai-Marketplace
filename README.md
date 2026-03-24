# TrustMint

**TrustMint** is a decentralized platform for training, verifying, and trading AI models with **Proof-of-Training (PoT)** guarantees. The platform ensures that AI models are authentic, verifiable, and securely shared in a marketplace, leveraging blockchain and IPFS technologies.

---

## 🌟 Project Vision

Many AI marketplaces face challenges like:

- **Fake/unauthenticated models**  
- **Lack of transparency in training data & process**  
- **Centralized control over model distribution**

**TrustMint solves this** by:

- Issuing cryptographic proofs of training for each AI model  
- Storing proofs and metadata on **IPFS**  
- Recording hash references on a **blockchain** for verification  
- Providing a **marketplace** where buyers can trust model authenticity  

---


**Components:**

1. **TrustMint CLI**  
   - Enables developers to train AI models securely  
   - Generates Proof-of-Training (PoT)  
   - Uploads model & proof to IPFS  

2. **Blockchain Smart Contract**  
   - Stores cryptographic hash of the model proof  
   - Allows buyers and developers to verify proof authenticity  

3. **IPFS / Decentralized Storage**  
   - Hosts model files & proof JSON  
   - Ensures content integrity & availability  

4. **Frontend Marketplace**  
   - Displays AI models for buyers  
   - Provides “Verify Proof” button (compares IPFS proof hash with blockchain)  
   - Guides for developers & buyers  

---

## ⚡ Features

### Developer Side

- Train AI models securely in Docker containers  
- Generate cryptographic **Proof-of-Training (PoT)**  
- Upload models & proofs to IPFS  
- View & manage model listings in dashboard  

### Buyer Side

- Browse AI models in the marketplace  
- View model metadata and training proof  
- Verify model authenticity using PoT  
- Purchase models using wallet integration  

### Core Principles

- **Decentralization:** Models and proofs stored on IPFS, verification on blockchain  
- **Transparency:** Buyers can verify model integrity before purchase  
- **Security:** Dockerized training and hashed proofs prevent tampering  

---

## 🖥 Frontend (Demo UI)

- Built with **React + Vite + Tailwind CSS**  
- Includes interactive pages:  
  - Landing / Hero section with **Download CLI button**  
  - Marketplace with **interactive model cards**  
  - Developer & Buyer Guides (step cards with badges)  
  - Model detail view (placeholder metadata & proof)  
  - Login page (MetaMask connect placeholder)  

> Note: Currently static demo; backend/blockchain integration can be added later.

---

## 🚀 Quick Start: Full System

### 1. Run the Backend Server
The backend server handles proof verification and artifact storage.

```bash
cd trustmint-backend
# Ensure dependencies are installed (e.g. Flask)
python3 app.py
```
*Server runs on http://localhost:5001*

### 2. Run the Training Workflow
Open a new terminal to run the training + verification flow.

```bash
cd trustmint-starter-kit

# Make scripts executable
chmod +x start-training.sh publish-artifacts.sh

# Step 1: Start the secure training environment
./start-training.sh
```

**Inside the Docker container:**
```bash
# Run the training command
trustmint train

# Exit the container when done
exit
```

**Back on your host machine:**
```bash
# Step 2: Publish and verify artifacts
./publish-artifacts.sh
```

---

## 🔧 Getting Started (Frontend Demo)

```bash
git clone https://github.com/<USERNAME>/trustmint-frontend.git
cd trustmint-frontend
npm install
npm run dev

Open http://localhost:5173
 in your browser.
```
## 🚧 Project Status
- [x] **CLI Tool**: Secure, Dockerized environment for training AI models.
- [x] **Proof Generation**: Generation of cryptographic Proof-of-Training (PoT).
- [x] **Backend**: Basic verification server and artifact separation.
- [x] **Frontend Demo**: Interactive UI for the marketplace.
- [ ] **Blockchain Integration**: Smart contracts for immutable proof storage.
- [ ] **Wallet Authentication**: Metamask integration for users.
- [ ] **Live Marketplace**: Full integration of uploads/verification with the frontend.


## 💡 Notes

The full TrustMint system consists of:

CLI for secure training

Blockchain smart contracts for proof storage

Decentralized file storage via IPFS


Buyer & developer marketplace