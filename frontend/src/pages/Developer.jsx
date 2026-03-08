import React, { useState, useEffect, useCallback } from 'react';
import { Link } from 'react-router-dom';
import { ethers } from 'ethers';
import styles from './Developer.module.css';

const BACKEND = import.meta.env.VITE_BACKEND_URL || 'http://localhost:5001';
const NFT_ADDRESS = import.meta.env.VITE_NFT_ADDRESS;
const MARKET_ADDRESS = import.meta.env.VITE_MARKETPLACE_ADDRESS;


const NFT_ABI = [
  'function totalSupply() view returns (uint256)',
  'function ownerOf(uint256 tokenId) view returns (address)',
  'function getModelMetadata(uint256 tokenId) view returns (string,string,bytes32,string,address,uint256,bool)',
  'function approve(address to, uint256 tokenId)',
  'function getApproved(uint256 tokenId) view returns (address)',
];

const MARKET_ABI = [
  'function listModel(address nftContract, uint256 tokenId, uint256 price) returns (uint256)',
  'function isListed(address nftContract, uint256 tokenId) view returns (bool)',
];

export default function Developer() {
  const [walletAddress, setWalletAddress] = useState('');
  const [isConnected, setIsConnected] = useState(false);
  const [isDownloading, setIsDownloading] = useState(false);
  const [myNFTs, setMyNFTs] = useState([]);
  const [nftsLoading, setNftsLoading] = useState(false);
  const [listingTokenId, setListingTokenId] = useState(null);

  useEffect(() => { checkWalletConnection(); }, []);

  const checkWalletConnection = async () => {
    if (window.ethereum) {
      try {
        const accounts = await window.ethereum.request({ method: 'eth_accounts' });
        if (accounts.length > 0) {
          setWalletAddress(accounts[0]);
          setIsConnected(true);
        }
      } catch (error) { console.error(error); }
    }
  };

  const connectWallet = async () => {
    if (!window.ethereum) {
      alert('MetaMask is not installed! Please install MetaMask to continue.');
      window.open('https://metamask.io/download/', '_blank');
      return;
    }
    try {
      const provider = new ethers.BrowserProvider(window.ethereum);
      const accounts = await provider.send("eth_requestAccounts", []);
      setWalletAddress(accounts[0]);
      setIsConnected(true);
      showToast('Wallet connected successfully!', 'success');
    } catch (error) {
      showToast('Failed to connect wallet', 'error');
    }
  };

  // Load NFTs owned by the connected wallet
  const loadMyNFTs = useCallback(async () => {
    if (!walletAddress) return;
    setNftsLoading(true);
    try {
      const provider = new ethers.BrowserProvider(window.ethereum);
      const nft = new ethers.Contract(NFT_ADDRESS, NFT_ABI, provider);
      const market = new ethers.Contract(MARKET_ADDRESS, MARKET_ABI, provider);
      const total = await nft.totalSupply();
      const owned = [];
      for (let i = 1; i <= Number(total); i++) {
        const owner = await nft.ownerOf(i);
        if (owner.toLowerCase() === walletAddress.toLowerCase()) {
          const meta = await nft.getModelMetadata(i);
          const listed = await market.isListed(NFT_ADDRESS, i);
          owned.push({
            tokenId: i,
            modelHash: meta[0],
            datasetHash: meta[1],
            ipfsCid: meta[3],
            listed,
          });
        }
      }
      setMyNFTs(owned);
    } catch (e) {
      console.error('Failed to load NFTs:', e);
    } finally {
      setNftsLoading(false);
    }
  }, [walletAddress]);

  useEffect(() => { if (isConnected) loadMyNFTs(); }, [isConnected, loadMyNFTs]);

  const listForSale = async (tokenId, priceEth) => {
    try {
      setListingTokenId(tokenId);
      const provider = new ethers.BrowserProvider(window.ethereum);
      const signer = await provider.getSigner();
      const nft = new ethers.Contract(NFT_ADDRESS, NFT_ABI, signer);
      const market = new ethers.Contract(MARKET_ADDRESS, MARKET_ABI, signer);
      const priceWei = ethers.parseEther(priceEth.toString());

      // Step 1: Approve marketplace to transfer the NFT
      showToast('Step 1/2: Approving marketplace in MetaMask...', 'info');
      const approveTx = await nft.approve(MARKET_ADDRESS, tokenId);
      await approveTx.wait();

      // Step 2: List on marketplace
      showToast('Step 2/2: Listing NFT on marketplace...', 'info');
      const listTx = await market.listModel(NFT_ADDRESS, tokenId, priceWei);
      await listTx.wait();

      showToast(`🎉 NFT #${tokenId} listed for ${priceEth} ETH!`, 'success');
      loadMyNFTs(); // Refresh
    } catch (e) {
      console.error(e);
      if (e.code === 'ACTION_REJECTED') showToast('Transaction rejected', 'error');
      else showToast(`Listing failed: ${e.message?.slice(0, 80)}`, 'error');
    } finally {
      setListingTokenId(null);
    }
  };

  const promptAndList = async (tokenId) => {
    const price = window.prompt(`Enter listing price in ETH for NFT #${tokenId}:`, '0.01');
    if (!price || isNaN(price) || Number(price) <= 0) return;
    await listForSale(tokenId, price);
  };

  const downloadCLI = async () => {
    if (!isConnected) { showToast('Please connect your wallet first!', 'error'); return; }
    try {
      setIsDownloading(true);
      const provider = new ethers.BrowserProvider(window.ethereum);
      const signer = await provider.getSigner();
      const message = `Bind Trustmint CLI to wallet: ${walletAddress}`;
      showToast('Please sign the message in MetaMask...', 'info');
      const signature = await signer.signMessage(message);
      const response = await fetch(`${BACKEND}/api/cli/download`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ walletAddress, signature, message })
      });
      if (!response.ok) throw new Error(`Server error: ${response.status}`);
      const blob = await response.blob();
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = 'trustmint';
      document.body.appendChild(a);
      a.click();
      document.body.removeChild(a);
      window.URL.revokeObjectURL(url);
      showToast('CLI downloaded successfully! Check your downloads folder.', 'success');
    } catch (error) {
      if (error.code === 'ACTION_REJECTED') showToast('Signature rejected', 'error');
      else showToast(`Download failed: ${error.message}`, 'error');
    } finally {
      setIsDownloading(false);
    }
  };

  const showToast = (message, type = 'success') => {
    const toast = document.createElement('div');
    toast.className = `toast ${type}`;
    toast.textContent = message;
    document.body.appendChild(toast);
    setTimeout(() => toast.classList.add('show'), 100);
    setTimeout(() => { toast.classList.remove('show'); setTimeout(() => toast.remove(), 300); }, 3500);
  };

  const shortAddress = (addr) => `${addr.slice(0, 6)}...${addr.slice(-4)}`;

  return (
    <main>
      <section className="page-section">
        <div className="container">
          <div className={styles.dashboardHeader}>
            <h2>Developer Dashboard</h2>
            <div style={{ display: 'flex', gap: '1rem', alignItems: 'center' }}>
              {isConnected ? (
                <>
                  <div style={{ padding: '0.5rem 1rem', background: 'var(--card)', borderRadius: 'var(--radius)', border: '1px solid var(--border)', display: 'flex', alignItems: 'center', gap: '0.5rem' }}>
                    <div style={{ width: '8px', height: '8px', borderRadius: '50%', background: '#10b981' }}></div>
                    <span style={{ fontWeight: '500' }}>{shortAddress(walletAddress)}</span>
                  </div>
                  <button className="btn btn-primary" onClick={downloadCLI} disabled={isDownloading}>
                    {isDownloading ? 'Downloading...' : '📥 Download CLI'}
                  </button>
                </>
              ) : (
                <button className="btn btn-primary" onClick={connectWallet}>🦊 Connect MetaMask</button>
              )}
            </div>
          </div>

          {!isConnected && (
            <div style={{ padding: '1rem', background: 'rgba(251,146,60,0.1)', border: '1px solid rgba(251,146,60,0.3)', borderRadius: 'var(--radius)', marginBottom: '2rem' }}>
              <strong>⚠️ Connect Your Wallet</strong>
              <p style={{ margin: '0.5rem 0 0 0', color: 'var(--text-secondary)' }}>
                Please connect your MetaMask wallet to download the CLI and publish models as NFTs.
              </p>
            </div>
          )}

          <h3>Quick Start Guide</h3>
          <div className={`${styles.steps} grid grid-5`}>
            {['Connect Wallet', 'Download CLI', 'Train Model', 'Publish as NFT', 'List & Earn'].map((step, i) => (
              <div className="card" key={i}>
                <h4>{i + 1}. {step}</h4>
                <p className="small-muted">{[
                  'Connect MetaMask to bind CLI to your address',
                  'Download wallet-bound Trustmint CLI tool',
                  'Run trustmint start to train your model',
                  'Run trustmint publish to mint NFT',
                  'Click "List for Sale" below to earn ETH',
                ][i]}</p>
              </div>
            ))}
          </div>

          <h3 style={{ marginTop: '2rem' }}>My NFTs</h3>
          <div className="card">
            {!isConnected ? (
              <p style={{ color: 'var(--text-secondary)', padding: '1rem' }}>Connect your wallet to see your minted NFTs.</p>
            ) : nftsLoading ? (
              <p style={{ color: 'var(--text-secondary)', padding: '1rem' }}>⏳ Loading your NFTs from blockchain...</p>
            ) : myNFTs.length === 0 ? (
              <div style={{ padding: '1.5rem', textAlign: 'center', color: 'var(--text-secondary)' }}>
                <p>🤖 No NFTs found for your wallet yet.</p>
                <p style={{ fontSize: '0.85rem' }}>Publish a model using the CLI to mint your first NFT.</p>
              </div>
            ) : (
              <table className="table">
                <thead>
                  <tr>
                    <th>Token ID</th>
                    <th>Model Hash</th>
                    <th>IPFS CID</th>
                    <th>Status</th>
                    <th>Actions</th>
                  </tr>
                </thead>
                <tbody>
                  {myNFTs.map(nft => (
                    <tr key={nft.tokenId}>
                      <td><strong>#{nft.tokenId}</strong></td>
                      <td style={{ fontSize: '0.75rem', wordBreak: 'break-all' }}>{nft.modelHash.slice(0, 16)}...</td>
                      <td style={{ fontSize: '0.75rem' }}>{nft.ipfsCid.slice(0, 16)}...</td>
                      <td>
                        {nft.listed
                          ? <span className={styles.statusPending}>🟢 Listed</span>
                          : <span className={styles.statusVerified}>🔵 Minted</span>}
                      </td>
                      <td>
                        {!nft.listed && (
                          <button
                            className="btn btn-primary"
                            style={{ padding: '0.3rem 0.7rem', fontSize: '0.8rem' }}
                            disabled={listingTokenId === nft.tokenId}
                            onClick={() => promptAndList(nft.tokenId)}
                          >
                            {listingTokenId === nft.tokenId ? '⏳ Listing...' : '💰 List for Sale'}
                          </button>
                        )}
                        {nft.listed && <span style={{ color: '#10b981', fontSize: '0.85rem' }}>✅ On Marketplace</span>}
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            )}
          </div>
        </div>
      </section>
    </main>
  );
}

