import React, { useState, useEffect } from 'react';
import { useParams, Link, useNavigate } from 'react-router-dom';
import { ethers } from 'ethers';

const BACKEND = import.meta.env.VITE_BACKEND_URL || 'http://localhost:5001';
const NFT_ADDRESS = import.meta.env.VITE_NFT_ADDRESS;
const MARKET_ADDRESS = import.meta.env.VITE_MARKETPLACE_ADDRESS;

const MARKET_ABI = [
    'function purchaseModel(uint256 listingId) payable',
    'function getListing(uint256 listingId) view returns (tuple(uint256 listingId, uint256 tokenId, address nftContract, address seller, uint256 price, bool active))',
];

const NFT_ABI = [
    'function getModelMetadata(uint256 tokenId) view returns (string,string,bytes32,string,address,uint256,bool)',
    'function tokenURI(uint256 tokenId) view returns (string)',
    'function ownerOf(uint256 tokenId) view returns (address)',
];

export default function ModelDetail() {
    const { tokenId } = useParams();
    const navigate = useNavigate();

    const [model, setModel] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [buying, setBuying] = useState(false);
    const [account, setAccount] = useState('');

    // Get current connected wallet
    useEffect(() => {
        if (window.ethereum) {
            window.ethereum.request({ method: 'eth_accounts' })
                .then(accs => { if (accs.length > 0) setAccount(accs[0].toLowerCase()); });
        }
    }, []);

    useEffect(() => {
        loadModel();
    }, [tokenId]);

    const loadModel = async () => {
        setLoading(true);
        setError(null);
        try {
            const provider = new ethers.BrowserProvider(window.ethereum);
            const nft = new ethers.Contract(NFT_ADDRESS, NFT_ABI, provider);
            const market = new ethers.Contract(MARKET_ADDRESS, MARKET_ABI, provider);

            // Fetch NFT metadata
            const meta = await nft.getModelMetadata(tokenId);
            const owner = await nft.ownerOf(tokenId);
            const tokenUri = await nft.tokenURI(tokenId);

            // Find active listing for this token
            // We get model list from backend and find matching token
            const res = await fetch(`${BACKEND}/api/models`);
            const models = await res.json();
            const listing = models.find(m => String(m.token_id) === String(tokenId));

            setModel({
                tokenId: Number(tokenId),
                modelHash: meta[0],
                datasetHash: meta[1],
                merkleRoot: meta[2],
                ipfsCid: meta[3],
                creator: meta[4],
                mintedAt: new Date(Number(meta[5]) * 1000).toLocaleString(),
                verified: meta[6],
                owner,
                tokenUri,
                listingId: listing?.listing_id,
                priceWei: listing?.price_wei,
                priceEth: listing?.price_eth,
                seller: listing?.seller,
                isListed: !!listing,
            });
        } catch (e) {
            console.error(e);
            setError('Failed to load model details. Make sure your wallet is connected to the local Hardhat network.');
        } finally {
            setLoading(false);
        }
    };

    const buyModel = async () => {
        if (!model?.listingId) return;
        try {
            setBuying(true);
            const provider = new ethers.BrowserProvider(window.ethereum);
            const accounts = await provider.send('eth_requestAccounts', []);
            const signer = await provider.getSigner();
            const market = new ethers.Contract(MARKET_ADDRESS, MARKET_ABI, signer);

            showToast('Confirm the purchase in MetaMask...', 'info');
            // Polygon Amoy minimum is ~25 gwei, overriding to avoid under-estimation errors
            const txOptions = {
                maxFeePerGas: ethers.parseUnits('30', 'gwei'),
                maxPriorityFeePerGas: ethers.parseUnits('30', 'gwei')
            };
            const tx = await market.purchaseModel(model.listingId, {
                value: model.priceWei,
                ...txOptions
            });
            showToast('⏳ Transaction submitted, waiting for confirmation...', 'info');
            await tx.wait();

            showToast(`🎉 NFT #${model.tokenId} purchased! You are now the owner.`, 'success');
            setAccount(accounts[0].toLowerCase());
            await loadModel(); // Refresh to show updated owner
        } catch (e) {
            console.error(e);
            if (e.code === 'ACTION_REJECTED') showToast('Transaction rejected.', 'error');
            else showToast(`Purchase failed: ${e.message?.slice(0, 100)}`, 'error');
        } finally {
            setBuying(false);
        }
    };

    const showToast = (message, type = 'success') => {
        const toast = document.createElement('div');
        toast.className = `toast ${type}`;
        toast.textContent = message;
        document.body.appendChild(toast);
        setTimeout(() => toast.classList.add('show'), 100);
        setTimeout(() => { toast.classList.remove('show'); setTimeout(() => toast.remove(), 300); }, 4000);
    };

    const shortAddr = (addr) => addr ? `${addr.slice(0, 6)}...${addr.slice(-4)}` : '—';
    const isOwner = account && model?.owner?.toLowerCase() === account;
    const isSeller = account && model?.seller?.toLowerCase() === account;

    const cardStyle = {
        background: 'var(--card)',
        border: '1px solid var(--border)',
        borderRadius: 'var(--radius)',
        padding: '1.5rem',
        marginBottom: '1.5rem',
    };

    const rowStyle = {
        display: 'flex',
        justifyContent: 'space-between',
        padding: '0.6rem 0',
        borderBottom: '1px solid var(--border)',
        gap: '1rem',
        flexWrap: 'wrap',
    };

    const hashStyle = {
        fontSize: '0.78rem',
        wordBreak: 'break-all',
        color: 'var(--text-secondary)',
        fontFamily: 'monospace',
    };

    if (loading) return (
        <main style={{ padding: '4rem', textAlign: 'center' }}>
            <p style={{ color: 'var(--text-secondary)', fontSize: '1.1rem' }}>
                ⏳ Loading model from blockchain...
            </p>
        </main>
    );

    if (error) return (
        <main style={{ padding: '4rem', textAlign: 'center' }}>
            <p style={{ color: '#ef4444' }}>⚠️ {error}</p>
            <Link to="/marketplace" className="btn btn-primary" style={{ marginTop: '1rem' }}>
                ← Back to Marketplace
            </Link>
        </main>
    );

    return (
        <main>
            <section className="page-section">
                <div className="container" style={{ maxWidth: '800px' }}>
                    <Link to="/marketplace" style={{ color: 'var(--text-secondary)', fontSize: '0.9rem' }}>
                        ← Back to Marketplace
                    </Link>

                    <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', margin: '1.5rem 0 1rem' }}>
                        <h2 style={{ margin: 0 }}>🤖 AI Model #{model.tokenId}</h2>
                        {model.verified && (
                            <span style={{ background: 'rgba(16,185,129,0.15)', color: '#10b981', padding: '0.3rem 0.8rem', borderRadius: '20px', fontSize: '0.85rem', fontWeight: '600' }}>
                                ✅ Cryptographically Verified
                            </span>
                        )}
                    </div>

                    {/* Buy / Price Section */}
                    {model.isListed ? (
                        <div style={{ ...cardStyle, background: 'rgba(99,102,241,0.08)', border: '1px solid rgba(99,102,241,0.3)' }}>
                            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', flexWrap: 'wrap', gap: '1rem' }}>
                                <div>
                                    <p style={{ margin: 0, color: 'var(--text-secondary)', fontSize: '0.85rem' }}>Listed price</p>
                                    <p style={{ margin: 0, fontSize: '2rem', fontWeight: '700' }}>{model.priceEth} ETH</p>
                                    <p style={{ margin: '0.25rem 0 0', color: 'var(--text-secondary)', fontSize: '0.84rem' }}>
                                        Seller: {shortAddr(model.seller)}
                                    </p>
                                </div>
                                <div style={{ display: 'flex', flexDirection: 'column', gap: '0.5rem', alignItems: 'flex-end' }}>
                                    {isOwner || isSeller ? (
                                        <div style={{ padding: '0.75rem 1.5rem', background: 'rgba(16,185,129,0.15)', color: '#10b981', borderRadius: 'var(--radius)', fontWeight: '600' }}>
                                            ✅ You own this NFT
                                        </div>
                                    ) : (
                                        <button
                                            className="btn btn-primary"
                                            style={{ padding: '0.75rem 2rem', fontSize: '1rem' }}
                                            onClick={buyModel}
                                            disabled={buying}
                                        >
                                            {buying ? '⏳ Processing...' : `🛒 Buy for ${model.priceEth} ETH`}
                                        </button>
                                    )}
                                    <p style={{ margin: 0, fontSize: '0.75rem', color: 'var(--text-secondary)' }}>
                                        +2.5% platform fee
                                    </p>
                                </div>
                            </div>
                        </div>
                    ) : (
                        <div style={{ ...cardStyle, background: 'rgba(251,146,60,0.08)', border: '1px solid rgba(251,146,60,0.3)' }}>
                            <p style={{ margin: 0, color: '#fb923c' }}>⚠️ This model is not currently listed for sale.</p>
                            {isOwner && <p style={{ margin: '0.5rem 0 0', fontSize: '0.85rem', color: 'var(--text-secondary)' }}>Go to Developer Dashboard to list it.</p>}
                        </div>
                    )}

                    {/* Ownership */}
                    <div style={cardStyle}>
                        <h4 style={{ margin: '0 0 1rem' }}>🔑 Ownership</h4>
                        <div style={rowStyle}>
                            <span style={{ fontWeight: '500' }}>Current Owner</span>
                            <span style={hashStyle}>{model.owner}</span>
                        </div>
                        <div style={rowStyle}>
                            <span style={{ fontWeight: '500' }}>Creator / Trainer</span>
                            <span style={hashStyle}>{model.creator}</span>
                        </div>
                        <div style={{ ...rowStyle, borderBottom: 'none' }}>
                            <span style={{ fontWeight: '500' }}>Minted At</span>
                            <span style={{ color: 'var(--text-secondary)' }}>{model.mintedAt}</span>
                        </div>
                    </div>

                    {/* Files — Owner only */}
                    {isOwner ? (
                        <div style={cardStyle}>
                            <h4 style={{ margin: '0 0 0.5rem' }}>📦 Your Files</h4>
                            <p style={{ margin: '0 0 1rem', color: 'var(--text-secondary)', fontSize: '0.88rem' }}>
                                As the owner, you have access to all model artifacts.
                            </p>
                            <div style={rowStyle}>
                                <span style={{ fontWeight: '500' }}>Model CID (IPFS)</span>
                                <span style={hashStyle}>{model.ipfsCid}</span>
                            </div>
                            <div style={{ ...rowStyle, borderBottom: 'none', alignItems: 'center' }}>
                                <div>
                                    <p style={{ margin: 0, fontWeight: '500' }}>Download Bundle</p>
                                    <p style={{ margin: '0.2rem 0 0', color: 'var(--text-secondary)', fontSize: '0.82rem' }}>
                                        Includes: model.pkl · train.py · trustmint.yml{model.includeDataset ? ' · dataset.zip' : ''}
                                    </p>
                                </div>
                                <button
                                    className="btn btn-primary"
                                    style={{ padding: '0.4rem 1rem', fontSize: '0.9rem', flexShrink: 0 }}
                                    onClick={async () => {
                                        try {
                                            showToast('⏳ Preparing your files...', 'info');
                                            const url = `${BACKEND}/api/artifacts/${model.modelHash}/download?wallet=${account}`;
                                            const res = await fetch(url);
                                            if (!res.ok) {
                                                const err = await res.json();
                                                throw new Error(err.error || res.statusText);
                                            }
                                            const blob = await res.blob();
                                            const a = document.createElement('a');
                                            a.href = window.URL.createObjectURL(blob);
                                            a.download = `trustmint-model-${model.tokenId}.zip`;
                                            document.body.appendChild(a);
                                            a.click();
                                            document.body.removeChild(a);
                                            showToast('✅ Download started!', 'success');
                                        } catch (e) {
                                            showToast('Download failed: ' + e.message, 'error');
                                        }
                                    }}
                                >
                                    ⬇️ Download All Files
                                </button>
                            </div>
                        </div>
                    ) : (
                        <div style={{ ...cardStyle, textAlign: 'center', padding: '2rem' }}>
                            <p style={{ fontSize: '1.5rem', margin: '0 0 0.5rem' }}>🔒</p>
                            <p style={{ fontWeight: '600', margin: '0 0 0.4rem' }}>Model files are for owners only</p>
                            <p style={{ color: 'var(--text-secondary)', fontSize: '0.88rem', margin: 0 }}>
                                Purchase this NFT to get access to model.pkl, train.py, trustmint.yml and optionally the dataset.
                            </p>
                        </div>
                    )}

                    {/* Cryptographic Proofs */}
                    <div style={cardStyle}>
                        <h4 style={{ margin: '0 0 1rem' }}>🔐 Cryptographic Proofs (On-Chain)</h4>
                        <div style={rowStyle}>
                            <span style={{ fontWeight: '500', minWidth: '120px' }}>Model Hash</span>
                            <span style={hashStyle}>{model.modelHash}</span>
                        </div>
                        <div style={rowStyle}>
                            <span style={{ fontWeight: '500', minWidth: '120px' }}>Dataset Hash</span>
                            <span style={hashStyle}>{model.datasetHash}</span>
                        </div>
                        <div style={{ ...rowStyle, borderBottom: 'none' }}>
                            <span style={{ fontWeight: '500', minWidth: '120px' }}>Merkle Root</span>
                            <span style={hashStyle}>{typeof model.merkleRoot === 'string' ? model.merkleRoot : ethers.hexlify(model.merkleRoot)}</span>
                        </div>
                    </div>

                    {/* How to verify */}
                    <div style={{ ...cardStyle, background: 'rgba(99,102,241,0.05)' }}>
                        <h4 style={{ margin: '0 0 0.75rem' }}>🧐 How to Verify Authenticity</h4>
                        <p style={{ margin: 0, color: 'var(--text-secondary)', fontSize: '0.9rem', lineHeight: '1.7' }}>
                            Download the model from IPFS using the CID above and run:
                            <code style={{ display: 'block', background: 'var(--bg)', padding: '0.5rem 0.75rem', borderRadius: '6px', margin: '0.5rem 0', fontSize: '0.82rem' }}>
                                sha256sum model.pkl
                            </code>
                            The output must match the <strong>Model Hash</strong> stored on-chain — proving the model was not tampered with after training.
                        </p>
                    </div>
                </div>
            </section>
        </main>
    );
}
