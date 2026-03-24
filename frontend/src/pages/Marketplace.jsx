import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import styles from './Marketplace.module.css';

const BACKEND = import.meta.env.VITE_BACKEND_URL || 'http://localhost:5001';

const ModelCard = ({ token_id, name, price_eth, creator, ipfs_cid, listing_id, verified }) => (
  <div className={styles.card}>
    <div className={styles.icon}>🤖</div>
    <div>
      <h4>{name}</h4>
      <p className="small-muted" style={{ wordBreak: 'break-all' }}>
        Creator: {creator ? `${creator.slice(0, 6)}...${creator.slice(-4)}` : 'Unknown'}
      </p>
      <p className="small-muted">
        {verified ? '✅ Verified' : '⏳ Unverified'}{' '}
        <span style={{ float: 'right' }}>Token #{token_id}</span>
      </p>
      {ipfs_cid && (
        <p className="small-muted" style={{ fontSize: '0.7rem', wordBreak: 'break-all' }}>
          IPFS: {ipfs_cid.slice(0, 20)}...
        </p>
      )}
      <div style={{ marginTop: 12 }}>
        <strong style={{ fontSize: 20 }}>{price_eth ? `${price_eth} ETH` : 'N/A'}</strong>
        <Link
          to={`/model/${token_id}`}
          className={styles.viewBtn}
        >
          View Details
        </Link>
      </div>
    </div>
  </div>
);

export default function Marketplace() {
  const [search, setSearch] = useState('');
  const [selectedCategory, setSelectedCategory] = useState('All');
  const [models, setModels] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const categories = ['All', 'NLP', 'Vision', 'Speech', 'Data Processing'];

  // Fetch live models from backend (reads from blockchain)
  useEffect(() => {
    const fetchModels = async () => {
      try {
        setLoading(true);
        const res = await fetch(`${BACKEND}/api/models`);
        if (!res.ok) throw new Error(`Server error: ${res.status}`);
        const data = await res.json();
        setModels(data);
        setError(null);
      } catch (err) {
        console.error('Failed to fetch models:', err);
        setError('Could not connect to backend. Is the server running?');
        setModels([]);
      } finally {
        setLoading(false);
      }
    };

    fetchModels();
    // Refresh every 15 seconds so new mints show up automatically
    const interval = setInterval(fetchModels, 15000);
    return () => clearInterval(interval);
  }, []);

  const filteredModels = models.filter((model) =>
    model.name.toLowerCase().includes(search.toLowerCase())
  );

  return (
    <main>
      <section className={`${styles.hero} page-section`}>
        <div className="container">
          <h2>AI Model Marketplace</h2>
          <p className="small-muted">
            Discover, evaluate, and purchase trusted AI models with verified provenance on the blockchain.
          </p>
          <div className={styles.searchRow}>
            <input
              placeholder="Search AI models..."
              value={search}
              onChange={(e) => setSearch(e.target.value)}
            />
            <select value={selectedCategory} onChange={(e) => setSelectedCategory(e.target.value)}>
              {categories.map((cat) => (
                <option key={cat} value={cat}>{cat}</option>
              ))}
            </select>
            <button className="btn btn-primary">Search</button>
          </div>
        </div>
      </section>

      <section className="container page-section">
        <div className={styles.filters}>
          {categories.map((cat) => (
            <button
              key={cat}
              className={`${styles.filterBtn} ${selectedCategory === cat ? styles.filterActive : ''}`}
              onClick={() => setSelectedCategory(cat)}
            >
              {cat}
            </button>
          ))}
        </div>
      </section>

      <section className="container page-section">
        {loading && (
          <p style={{ textAlign: 'center', color: 'var(--text-secondary)' }}>
            ⏳ Loading models from blockchain...
          </p>
        )}

        {error && (
          <div style={{
            padding: '1rem',
            background: 'rgba(239,68,68,0.1)',
            border: '1px solid rgba(239,68,68,0.3)',
            borderRadius: 'var(--radius)',
            marginBottom: '1rem'
          }}>
            ⚠️ {error}
          </div>
        )}

        {!loading && !error && filteredModels.length === 0 && (
          <div style={{ textAlign: 'center', padding: '3rem', color: 'var(--text-secondary)' }}>
            <p style={{ fontSize: '3rem' }}>🤖</p>
            <h3>No models listed yet</h3>
            <p>Train and publish your first model using the Trustmint CLI!</p>
            <Link to="/developer" className="btn btn-primary" style={{ marginTop: '1rem' }}>
              Go to Developer Dashboard
            </Link>
          </div>
        )}

        {!loading && filteredModels.length > 0 && (
          <>
            <h3>Available Models ({filteredModels.length})</h3>
            <div className={`${styles.grid} grid grid-4`}>
              {filteredModels.map((model) => (
                <ModelCard key={model.listing_id} {...model} />
              ))}
            </div>
          </>
        )}
      </section>
    </main>
  );
}