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

## 🔧 Getting Started (Frontend Demo)

```bash
git clone https://github.com/<USERNAME>/trustmint-frontend.git
cd trustmint-frontend
npm install
npm run dev

Open http://localhost:5173
 in your browser.
```
## 🛠 Future Development

Integrate real blockchain for proof storage

Implement wallet-based authentication & payments

Add real model uploads & verification logic

Improve UX with live data, charts, and analytics


## 💡 Notes

This repository currently contains the frontend demo UI. The full TrustMint system consists of:

CLI for secure training

Blockchain smart contracts for proof storage

Decentralized file storage via IPFS


Buyer & developer marketplace