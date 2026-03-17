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
            merkleRoot: meta[2],
            ipfsCid: meta[3],
            creator: meta[4],
            mintedAt: new Date(Number(meta[5]) * 1000).toLocaleString(),
            verified: meta[6],
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

      // Polygon Amoy requires a minimum of 25 gwei, but MetaMask frequently under-estimates it.
      const txOptions = {
        maxFeePerGas: ethers.parseUnits('30', 'gwei'),
        maxPriorityFeePerGas: ethers.parseUnits('30', 'gwei')
      };

      // Step 1: Approve marketplace to transfer the NFT
      showToast('Step 1/2: Approving marketplace in MetaMask...', 'info');
      const approveTx = await nft.approve(MARKET_ADDRESS, tokenId, txOptions);
      await approveTx.wait();

      // Step 2: List on marketplace
      showToast('Step 2/2: Listing NFT on marketplace...', 'info');
      const listTx = await market.listModel(NFT_ADDRESS, tokenId, priceWei, txOptions);
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

  const promptAndList = async (tokenId, modelHash) => {
    const price = window.prompt(`Enter listing price in ETH for NFT #${tokenId}:`, '0.01');
    if (!price || isNaN(price) || Number(price) <= 0) return;

    const includeDataset = window.confirm(
      `Include dataset with purchase?\n\n` +
      `YES → Buyers who purchase this NFT can also download your training dataset.\n` +
      `NO  → Buyers only get the model file, training script, and config.`
    );

    // Save dataset preference to backend
    if (modelHash) {
      try {
        await fetch(`${BACKEND}/api/listing-prefs/${modelHash}`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ include_dataset: includeDataset }),
        });
      } catch (e) { console.warn('Could not save listing prefs:', e); }
    }

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
              myNFTs.map(nft => {
                const rowStyle = { display: 'flex', justifyContent: 'space-between', padding: '0.45rem 0', borderBottom: '1px solid var(--border)', gap: '1rem', flexWrap: 'wrap' };
                const monoStyle = { fontSize: '0.74rem', wordBreak: 'break-all', color: 'var(--text-secondary)', fontFamily: 'monospace' };
                return (
                  <div key={nft.tokenId} style={{ border: '1px solid var(--border)', borderRadius: 'var(--radius)', padding: '1.25rem', marginBottom: '1rem' }}>
                    {/* Header */}
                    <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '1rem', flexWrap: 'wrap', gap: '0.5rem' }}>
                      <div style={{ display: 'flex', alignItems: 'center', gap: '0.6rem' }}>
                        <span style={{ fontSize: '1.3rem' }}>🤖</span>
                        <strong style={{ fontSize: '1rem' }}>AI Model #{nft.tokenId}</strong>
                        {nft.verified && <span style={{ color: '#10b981', fontSize: '0.8rem' }}>✅ Verified</span>}
                      </div>
                      <div style={{ display: 'flex', gap: '0.5rem', alignItems: 'center', flexWrap: 'wrap' }}>
                        {nft.listed
                          ? <span className={styles.statusPending}>🟢 Listed on Marketplace</span>
                          : <span className={styles.statusVerified}>🔵 Minted</span>}
                        <Link to={`/model/${nft.tokenId}`} className="btn btn-secondary" style={{ padding: '0.25rem 0.65rem', fontSize: '0.8rem' }}>View Page</Link>
                        {!nft.listed && (
                          <button
                            className="btn btn-primary"
                            style={{ padding: '0.25rem 0.7rem', fontSize: '0.8rem' }}
                            disabled={listingTokenId === nft.tokenId}
                            onClick={() => promptAndList(nft.tokenId, nft.modelHash)}
                          >
                            {listingTokenId === nft.tokenId ? '⏳ Listing...' : '💰 List for Sale'}
                          </button>
                        )}
                      </div>
                    </div>

                    {/* Details */}
                    <div style={rowStyle}><span style={{ fontWeight: '500', minWidth: 120 }}>Minted At</span><span style={{ color: 'var(--text-secondary)', fontSize: '0.85rem' }}>{nft.mintedAt}</span></div>
                    <div style={rowStyle}><span style={{ fontWeight: '500', minWidth: 120 }}>Model Hash</span><span style={monoStyle}>{nft.modelHash}</span></div>
                    <div style={rowStyle}><span style={{ fontWeight: '500', minWidth: 120 }}>Dataset Hash</span><span style={monoStyle}>{nft.datasetHash}</span></div>
                    <div style={rowStyle}><span style={{ fontWeight: '500', minWidth: 120 }}>IPFS CID</span><span style={monoStyle}>{nft.ipfsCid}</span></div>

                    {/* Download */}
                    <div style={{ ...rowStyle, borderBottom: 'none', alignItems: 'center', marginTop: '0.5rem' }}>
                      <div>
                        <span style={{ fontWeight: '500' }}>Download Bundle</span>
                        <span style={{ marginLeft: '0.5rem', color: 'var(--text-secondary)', fontSize: '0.78rem' }}>model.pkl · train.py · trustmint.yml</span>
                      </div>
                      <button
                        className="btn btn-primary"
                        style={{ padding: '0.3rem 0.8rem', fontSize: '0.82rem' }}
                        onClick={async () => {
                          try {
                            showToast('⏳ Preparing download...', 'info');
                            const url = `${BACKEND}/api/artifacts/${nft.modelHash}/download?wallet=${walletAddress}`;
                            const res = await fetch(url);
                            if (!res.ok) { const e = await res.json(); throw new Error(e.error); }
                            const blob = await res.blob();
                            const a = document.createElement('a');
                            a.href = window.URL.createObjectURL(blob);
                            a.download = `trustmint-model-${nft.tokenId}.zip`;
                            a.click();
                            showToast('✅ Download started!', 'success');
                          } catch (e) { showToast('Download failed: ' + e.message, 'error'); }
                        }}
                      >
                        ⬇️ Download All Files
                      </button>
                    </div>
                  </div>
                );
              })
            )}
          </div>
        </div>
      </section>
    </main>
  );
}

