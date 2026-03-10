import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { ethers } from 'ethers';
import { auth } from '../firebase';
import { onAuthStateChanged } from 'firebase/auth';

const BACKEND = import.meta.env.VITE_BACKEND_URL || 'http://localhost:5001';
const NFT_ADDRESS = import.meta.env.VITE_NFT_ADDRESS;
const MARKET_ADDRESS = import.meta.env.VITE_MARKETPLACE_ADDRESS;
const RPC_URL = 'http://127.0.0.1:8545'; // Direct node — avoids MetaMask eth_getLogs issues

const NFT_ABI = [
  'function totalSupply() view returns (uint256)',
  'function ownerOf(uint256 tokenId) view returns (address)',
  'function getModelMetadata(uint256 tokenId) view returns (string,string,bytes32,string,address,uint256,bool)',
  'function tokenURI(uint256 tokenId) view returns (string)',
  'event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)',
];

const MARKET_ABI = [
  'function isListed(address nftContract, uint256 tokenId) view returns (bool)',
  'event ListingCreated(uint256 indexed listingId, uint256 indexed tokenId, address indexed seller, uint256 price)',
  'event ModelPurchased(uint256 indexed listingId, uint256 indexed tokenId, address seller, address buyer, uint256 price)',
];

const BADGE = {
  minted: { bg: 'rgba(99,102,241,0.15)', color: '#818cf8', label: '🪙 Minted' },
  bought: { bg: 'rgba(16,185,129,0.15)', color: '#10b981', label: '🛒 Purchased' },
  sold: { bg: 'rgba(251,146,60,0.15)', color: '#fb923c', label: '💸 Sold' },
  listed: { bg: 'rgba(234,179,8,0.15)', color: '#eab308', label: '🏷️ Listed' },
  received: { bg: 'rgba(14,165,233,0.15)', color: '#38bdf8', label: '📥 Received' },
};
function StatusBadge({ type }) {
  const b = BADGE[type] || BADGE.minted;
  return <span style={{ background: b.bg, color: b.color, padding: '0.2rem 0.6rem', borderRadius: '20px', fontSize: '0.75rem', fontWeight: '600', whiteSpace: 'nowrap' }}>{b.label}</span>;
}

export default function Profile() {
  const [user, setUser] = useState(null);
  const [fbLoading, setFbLoading] = useState(true);
  const [wallet, setWallet] = useState('');
  const [tab, setTab] = useState('owned');
  const [history, setHistory] = useState([]);
  const [ownedNFTs, setOwnedNFTs] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    const unsub = onAuthStateChanged(auth, (u) => {
      if (!u) navigate('/login');
      else { setUser(u); setFbLoading(false); }
    });
    return unsub;
  }, [navigate]);

  useEffect(() => {
    if (window.ethereum) {
      window.ethereum.request({ method: 'eth_accounts' })
        .then(accs => { if (accs[0]) setWallet(accs[0].toLowerCase()); });
      window.ethereum.on('accountsChanged', (accs) => setWallet(accs[0]?.toLowerCase() || ''));
    }
  }, []);

  useEffect(() => { if (wallet) loadData(); }, [wallet]);

  const connectWallet = async () => {
    try {
      const provider = new ethers.BrowserProvider(window.ethereum);
      const accs = await provider.send('eth_requestAccounts', []);
      setWallet(accs[0].toLowerCase());
    } catch (e) { setError('Could not connect wallet.'); }
  };

  const loadData = async () => {
    if (!wallet || !NFT_ADDRESS) return;
    setLoading(true);
    setError('');
    try {
      // Use direct RPC provider for event queries (MetaMask blocks eth_getLogs)
      const rpc = new ethers.JsonRpcProvider(RPC_URL);
      const nft = new ethers.Contract(NFT_ADDRESS, NFT_ABI, rpc);
      const market = new ethers.Contract(MARKET_ADDRESS, MARKET_ABI, rpc);

      const walletCS = ethers.getAddress(wallet);
      const events = [];

      // --- Events ---
      // Mints (Transfer from 0x0 to wallet)
      const mints = await nft.queryFilter(nft.filters.Transfer(ethers.ZeroAddress, walletCS), 0, 'latest');
      for (const e of mints) {
        const blk = await rpc.getBlock(e.blockNumber);
        events.push({ type: 'minted', tokenId: Number(e.args.tokenId), txHash: e.transactionHash, time: blk.timestamp });
      }

      // Received (Transfer from non-0 to wallet)
      const recvs = await nft.queryFilter(nft.filters.Transfer(null, walletCS), 0, 'latest');
      for (const e of recvs) {
        if (e.args.from === ethers.ZeroAddress) continue;
        const blk = await rpc.getBlock(e.blockNumber);
        events.push({ type: 'received', tokenId: Number(e.args.tokenId), from: e.args.from, txHash: e.transactionHash, time: blk.timestamp });
      }

      // Sold (Transfer from wallet to non-0)
      const sents = await nft.queryFilter(nft.filters.Transfer(walletCS, null), 0, 'latest');
      for (const e of sents) {
        if (e.args.to === ethers.ZeroAddress) continue;
        const blk = await rpc.getBlock(e.blockNumber);
        events.push({ type: 'sold', tokenId: Number(e.args.tokenId), to: e.args.to, txHash: e.transactionHash, time: blk.timestamp });
      }

      // Listed
      const lists = await market.queryFilter(market.filters.ListingCreated(null, null, walletCS), 0, 'latest');
      for (const e of lists) {
        const blk = await rpc.getBlock(e.blockNumber);
        events.push({ type: 'listed', tokenId: Number(e.args.tokenId), price: ethers.formatEther(e.args.price), txHash: e.transactionHash, time: blk.timestamp });
      }

      // Purchases (buyer = wallet)
      const buys = await market.queryFilter(market.filters.ModelPurchased(), 0, 'latest');
      for (const e of buys) {
        if (e.args.buyer?.toLowerCase() !== wallet) continue;
        const blk = await rpc.getBlock(e.blockNumber);
        events.push({ type: 'bought', tokenId: Number(e.args.tokenId), price: ethers.formatEther(e.args.price), from: e.args.seller, txHash: e.transactionHash, time: blk.timestamp });
      }

      events.sort((a, b) => b.time - a.time);
      setHistory(events);

      // --- Currently owned ---
      const total = Number(await nft.totalSupply());
      const owned = [];
      for (let i = 1; i <= total; i++) {
        const owner = await nft.ownerOf(i);
        if (owner.toLowerCase() !== wallet) continue;
        const meta = await nft.getModelMetadata(i);
        const uri = await nft.tokenURI(i);
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
          tokenUri: uri,
          listed,
        });
      }
      setOwnedNFTs(owned);
    } catch (e) {
      console.error(e);
      setError('Failed to load data. Make sure you are connected to the Hardhat network.');
    } finally {
      setLoading(false);
    }
  };

  const shortAddr = a => a ? `${a.slice(0, 6)}...${a.slice(-4)}` : '—';
  const shortHash = h => h ? `${h.slice(0, 10)}...` : '—';
  const timeStr = ts => new Date(ts * 1000).toLocaleString();

  const card = { background: 'var(--card)', border: '1px solid var(--border)', borderRadius: 'var(--radius)', padding: '1.5rem', marginBottom: '1.5rem' };
  const row = { display: 'flex', justifyContent: 'space-between', padding: '0.5rem 0', borderBottom: '1px solid var(--border)', gap: '1rem', flexWrap: 'wrap' };
  const mono = { fontSize: '0.75rem', wordBreak: 'break-all', color: 'var(--text-secondary)', fontFamily: 'monospace' };

  const TabBtn = ({ id, label }) => (
    <button onClick={() => setTab(id)} style={{ padding: '0.45rem 1.1rem', borderRadius: 'var(--radius)', border: '1px solid var(--border)', background: tab === id ? 'var(--primary)' : 'var(--card)', color: tab === id ? '#fff' : 'var(--text)', cursor: 'pointer', fontWeight: '500', fontSize: '0.9rem' }}>
      {label}
    </button>
  );

  if (fbLoading) return <main style={{ padding: '4rem', textAlign: 'center' }}>Loading...</main>;

  return (
    <main>
      <section className="page-section">
        <div className="container" style={{ maxWidth: '900px' }}>

          {/* Account card */}
          <div style={{ ...card, display: 'flex', justifyContent: 'space-between', alignItems: 'center', flexWrap: 'wrap', gap: '1rem' }}>
            <div>
              <h2 style={{ margin: '0 0 0.25rem' }}>👤 {user?.displayName || user?.email?.split('@')[0]}</h2>
              <p style={{ margin: 0, color: 'var(--text-secondary)', fontSize: '0.88rem' }}>{user?.email}</p>
              <p style={{ margin: '0.2rem 0 0', color: 'var(--text-secondary)', fontSize: '0.8rem' }}>
                Joined: {user?.metadata?.creationTime?.split(' ').slice(0, 4).join(' ')}
              </p>
            </div>
            {wallet ? (
              <div style={{ display: 'flex', alignItems: 'center', gap: '0.5rem', background: 'rgba(16,185,129,0.1)', padding: '0.5rem 0.9rem', borderRadius: 'var(--radius)', border: '1px solid rgba(16,185,129,0.3)' }}>
                <div style={{ width: 8, height: 8, borderRadius: '50%', background: '#10b981' }}></div>
                <span style={{ fontFamily: 'monospace', fontSize: '0.85rem' }}>{shortAddr(wallet)}</span>
              </div>
            ) : (
              <button className="btn btn-primary" onClick={connectWallet}>🦊 Connect Wallet</button>
            )}
          </div>

          {error && <p style={{ color: '#ef4444', marginBottom: '1rem' }}>⚠️ {error}</p>}

          {!wallet ? (
            <div style={{ ...card, textAlign: 'center', padding: '3rem' }}>
              <p style={{ fontSize: '2rem' }}>🦊</p>
              <p style={{ fontWeight: '600', marginBottom: '1rem' }}>Connect your wallet to see your NFT history</p>
              <button className="btn btn-primary" onClick={connectWallet}>Connect MetaMask</button>
            </div>
          ) : (
            <>
              <div style={{ display: 'flex', gap: '0.5rem', marginBottom: '1.5rem', flexWrap: 'wrap' }}>
                <TabBtn id="owned" label={`🤖 My NFTs (${ownedNFTs.length})`} />
                <TabBtn id="history" label={`📜 History (${history.length})`} />
              </div>

              {loading && (
                <div style={{ textAlign: 'center', padding: '3rem', color: 'var(--text-secondary)' }}>⏳ Reading blockchain...</div>
              )}

              {/* ── OWNED NFTs Tab ── */}
              {!loading && tab === 'owned' && (
                ownedNFTs.length === 0 ? (
                  <div style={{ ...card, textAlign: 'center', padding: '2.5rem' }}>
                    <p style={{ fontSize: '2rem' }}>🤖</p>
                    <p>You don't own any NFTs currently.</p>
                    <Link to="/marketplace" className="btn btn-primary" style={{ marginTop: '0.75rem' }}>Browse Marketplace</Link>
                  </div>
                ) : (
                  ownedNFTs.map(nft => (
                    <div key={nft.tokenId} style={{ ...card, marginBottom: '1rem' }}>
                      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '1rem', flexWrap: 'wrap', gap: '0.5rem' }}>
                        <div style={{ display: 'flex', alignItems: 'center', gap: '0.75rem' }}>
                          <span style={{ fontSize: '1.4rem' }}>🤖</span>
                          <div>
                            <strong style={{ fontSize: '1.05rem' }}>AI Model #{nft.tokenId}</strong>
                            {nft.verified && <span style={{ marginLeft: '0.5rem', color: '#10b981', fontSize: '0.8rem' }}>✅ Verified</span>}
                          </div>
                        </div>
                        <div style={{ display: 'flex', gap: '0.5rem' }}>
                          <StatusBadge type={nft.listed ? 'listed' : 'minted'} />
                          <Link to={`/model/${nft.tokenId}`} className="btn btn-secondary" style={{ padding: '0.25rem 0.65rem', fontSize: '0.8rem' }}>View Page</Link>
                          <Link to="/developer" className="btn btn-primary" style={{ padding: '0.25rem 0.65rem', fontSize: '0.8rem' }}>{nft.listed ? 'Manage' : '💰 List for Sale'}</Link>
                        </div>
                      </div>

                      {/* Full model details */}
                      <div style={row}><span style={{ fontWeight: '500' }}>Minted At</span><span style={{ color: 'var(--text-secondary)' }}>{nft.mintedAt}</span></div>
                      <div style={row}><span style={{ fontWeight: '500' }}>Creator</span><span style={mono}>{nft.creator}</span></div>
                      <div style={row}><span style={{ fontWeight: '500' }}>Model Hash</span><span style={mono}>{nft.modelHash}</span></div>
                      <div style={row}><span style={{ fontWeight: '500' }}>Dataset Hash</span><span style={mono}>{nft.datasetHash}</span></div>
                      <div style={row}><span style={{ fontWeight: '500' }}>IPFS CID</span><span style={mono}>{nft.ipfsCid}</span></div>
                      <div style={{ ...row, borderBottom: 'none', alignItems: 'center' }}>
                        <div>
                          <span style={{ fontWeight: '500' }}>Download Bundle</span>
                          <span style={{ marginLeft: '0.5rem', color: 'var(--text-secondary)', fontSize: '0.8rem' }}>model.pkl · train.py · trustmint.yml</span>
                        </div>
                        <button
                          className="btn btn-primary"
                          style={{ padding: '0.3rem 0.8rem', fontSize: '0.82rem' }}
                          onClick={async () => {
                            try {
                              const url = `${BACKEND}/api/artifacts/${nft.modelHash}/download?wallet=${wallet}`;
                              const res = await fetch(url);
                              if (!res.ok) { const e = await res.json(); throw new Error(e.error); }
                              const blob = await res.blob();
                              const a = document.createElement('a');
                              a.href = window.URL.createObjectURL(blob);
                              a.download = `trustmint-model-${nft.tokenId}.zip`;
                              a.click();
                            } catch (e) { alert('Download failed: ' + e.message); }
                          }}
                        >⬇️ Download All</button>
                      </div>
                    </div>
                  ))
                )
              )}

              {/* ── HISTORY Tab ── */}
              {!loading && tab === 'history' && (
                history.length === 0 ? (
                  <div style={{ ...card, textAlign: 'center', padding: '2.5rem', color: 'var(--text-secondary)' }}>
                    <p style={{ fontSize: '2rem' }}>📭</p>
                    <p>No transactions found for this wallet on the current network.</p>
                  </div>
                ) : (
                  <div style={card}>
                    <table className="table">
                      <thead><tr><th>Type</th><th>Token</th><th>Details</th><th>Time</th><th>Tx Hash</th></tr></thead>
                      <tbody>
                        {history.map((e, i) => (
                          <tr key={i}>
                            <td><StatusBadge type={e.type} /></td>
                            <td><Link to={`/model/${e.tokenId}`} style={{ color: 'var(--primary)', fontWeight: '600' }}>#{e.tokenId}</Link></td>
                            <td style={{ fontSize: '0.82rem', color: 'var(--text-secondary)' }}>
                              {e.type === 'minted' && 'Token created'}
                              {e.type === 'listed' && `Listed for ${e.price} ETH`}
                              {e.type === 'bought' && `Bought for ${e.price} ETH from ${shortAddr(e.from)}`}
                              {e.type === 'sold' && `Transferred to ${shortAddr(e.to)}`}
                              {e.type === 'received' && `Received from ${shortAddr(e.from)}`}
                            </td>
                            <td style={{ fontSize: '0.78rem', color: 'var(--text-secondary)', whiteSpace: 'nowrap' }}>{timeStr(e.time)}</td>
                            <td><span title={e.txHash} style={mono}>{shortHash(e.txHash)}</span></td>
                          </tr>
                        ))}
                      </tbody>
                    </table>
                  </div>
                )
              )}
            </>
          )}
        </div>
      </section>
    </main>
  );
}