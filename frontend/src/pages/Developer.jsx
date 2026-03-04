import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { ethers } from 'ethers';
import styles from './Developer.module.css';

export default function Developer() {
  const [walletAddress, setWalletAddress] = useState('');
  const [isConnected, setIsConnected] = useState(false);
  const [isDownloading, setIsDownloading] = useState(false);

  // Check if wallet is already connected
  useEffect(() => {
    checkWalletConnection();
  }, []);

  const checkWalletConnection = async () => {
    if (window.ethereum) {
      try {
        const accounts = await window.ethereum.request({ method: 'eth_accounts' });
        if (accounts.length > 0) {
          setWalletAddress(accounts[0]);
          setIsConnected(true);
        }
      } catch (error) {
        console.error('Error checking wallet connection:', error);
      }
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
      console.error('Error connecting wallet:', error);
      showToast('Failed to connect wallet', 'error');
    }
  };

  const downloadCLI = async () => {
    if (!isConnected) {
      showToast('Please connect your wallet first!', 'error');
      return;
    }

    try {
      setIsDownloading(true);

      // Request signature to prove wallet ownership
      const provider = new ethers.BrowserProvider(window.ethereum);
      const signer = await provider.getSigner();
      const message = `Bind Trustmint CLI to wallet: ${walletAddress}`;

      showToast('Please sign the message in MetaMask...', 'info');
      const signature = await signer.signMessage(message);

      // Call backend API
      const response = await fetch('http://localhost:5001/api/cli/download', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          walletAddress,
          signature,
          message
        })
      });

      if (!response.ok) {
        throw new Error(`Server error: ${response.status}`);
      }

      // Download the CLI binary
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
      console.error('Error downloading CLI:', error);
      if (error.code === 'ACTION_REJECTED') {
        showToast('Signature rejected', 'error');
      } else {
        showToast(`Download failed: ${error.message}`, 'error');
      }
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
    setTimeout(() => {
      toast.classList.remove('show');
      setTimeout(() => toast.remove(), 300);
    }, 3000);
  };

  const shortAddress = (addr) => {
    return `${addr.slice(0, 6)}...${addr.slice(-4)}`;
  };

  return (
    <main>
      <section className="page-section">
        <div className="container">
          <div className={styles.dashboardHeader}>
            <h2>Developer Dashboard</h2>
            <div style={{ display: 'flex', gap: '1rem', alignItems: 'center' }}>
              {isConnected ? (
                <>
                  <div style={{
                    padding: '0.5rem 1rem',
                    background: 'var(--card)',
                    borderRadius: 'var(--radius)',
                    border: '1px solid var(--border)',
                    display: 'flex',
                    alignItems: 'center',
                    gap: '0.5rem'
                  }}>
                    <div style={{
                      width: '8px',
                      height: '8px',
                      borderRadius: '50%',
                      background: '#10b981'
                    }}></div>
                    <span style={{ fontWeight: '500' }}>{shortAddress(walletAddress)}</span>
                  </div>
                  <button
                    className="btn btn-primary"
                    onClick={downloadCLI}
                    disabled={isDownloading}
                  >
                    {isDownloading ? 'Downloading...' : '📥 Download CLI'}
                  </button>
                </>
              ) : (
                <button className="btn btn-primary" onClick={connectWallet}>
                  🦊 Connect MetaMask
                </button>
              )}
            </div>
          </div>

          {!isConnected && (
            <div style={{
              padding: '1rem',
              background: 'rgba(251, 146, 60, 0.1)',
              border: '1px solid rgba(251, 146, 60, 0.3)',
              borderRadius: 'var(--radius)',
              marginBottom: '2rem'
            }}>
              <strong>⚠️ Connect Your Wallet</strong>
              <p style={{ margin: '0.5rem 0 0 0', color: 'var(--text-secondary)' }}>
                Please connect your MetaMask wallet to download the CLI and publish models as NFTs.
              </p>
            </div>
          )}

          <h3>Quick Start Guide</h3>
          <div className={`${styles.steps} grid grid-5`}>
            <div className="card">
              <h4>1. Connect Wallet</h4>
              <p className="small-muted">Connect MetaMask to bind CLI to your address</p>
            </div>
            <div className="card">
              <h4>2. Download CLI</h4>
              <p className="small-muted">Download wallet-bound Trustmint CLI tool</p>
            </div>
            <div className="card">
              <h4>3. Train Model</h4>
              <p className="small-muted">Run <code style={{ background: 'var(--bg)', padding: '0.1rem 0.4rem', borderRadius: '4px' }}>trustmint train</code> to train your model</p>
            </div>
            <div className="card">
              <h4>4. Publish as NFT</h4>
              <p className="small-muted">Run <code style={{ background: 'var(--bg)', padding: '0.1rem 0.4rem', borderRadius: '4px' }}>trustmint publish</code> to mint NFT</p>
            </div>
            <div className="card">
              <h4>5. Start Earning</h4>
              <p className="small-muted">Your NFT is listed on the marketplace automatically</p>
            </div>
          </div>

          <h3>Earnings Overview</h3>
          <div className={`${styles.earnings} grid grid-4`}>
            <div className="card">
              <p className="small-muted">This Month</p>
              <h4>$2,450.00</h4>
            </div>
            <div className="card">
              <p className="small-muted">Last Month</p>
              <h4>$1,890.00</h4>
            </div>
            <div className="card">
              <p className="small-muted">Total Earned</p>
              <h4>$12,340.00</h4>
            </div>
            <div className="card">
              <p className="small-muted">Pending Payout</p>
              <h4>$340.00</h4>
            </div>
          </div>

          <h3>Your NFT Models</h3>
          <div className="card">
            <div className={styles.tableHeader}>
              <Link to="/upload" className="btn btn-secondary">Upload New</Link>
            </div>
            <table className="table">
              <thead>
                <tr>
                  <th>Model Name</th>
                  <th>Token ID</th>
                  <th>Status</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>ImageNet Classifier</td>
                  <td>#42</td>
                  <td><span className={styles.statusVerified}>Minted</span></td>
                  <td><button className="btn btn-primary" style={{ padding: '0.25rem 0.5rem', fontSize: '0.75rem' }}>⋮</button></td>
                </tr>
                <tr>
                  <td>Text Summarizer</td>
                  <td>#43</td>
                  <td><span className={styles.statusPending}>Listed</span></td>
                  <td><button className="btn btn-primary" style={{ padding: '0.25rem 0.5rem', fontSize: '0.75rem' }}>⋮</button></td>
                </tr>
                <tr>
                  <td>Speech2Text</td>
                  <td>#44</td>
                  <td><span className={styles.statusVerified}>Minted</span></td>
                  <td><button className="btn btn-primary" style={{ padding: '0.25rem 0.5rem', fontSize: '0.75rem' }}>⋮</button></td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </section>
    </main>
  );
}