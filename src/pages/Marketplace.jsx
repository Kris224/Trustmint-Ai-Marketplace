import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import styles from './Marketplace.module.css';

const ModelCard = ({ name, price, score, category }) => (
  <div className={styles.card}>
    <div className={styles.icon}></div>
    <div>
      <h4>{name}</h4>
      <p className="small-muted">
        Trust Score: {score} <span style={{ float: 'right' }}>★ 4.7</span>
      </p>
      <div style={{ marginTop: 12 }}>
        <strong style={{ fontSize: 20 }}>${price}</strong>
        <Link to={`/model/${name.toLowerCase().replace(/\s+/g, '-')}`} className={styles.viewBtn}>
          View Details
        </Link>
      </div>
    </div>
  </div>
);

export default function Marketplace() {
  const [search, setSearch] = useState('');
  const [selectedCategory, setSelectedCategory] = useState('All');
  const categories = ['All', 'NLP', 'Vision', 'Speech', 'Data Processing'];

  const rawModels = [
    { name: 'TextGenX', price: 49, score: 92, category: 'NLP' },
    { name: 'VisionPro', price: 39, score: 88, category: 'Vision' },
    { name: 'SpeechAI Lite', price: 59, score: 95, category: 'Speech' },
    { name: 'DataSummarizer', price: 29, score: 85, category: 'Data Processing' },
  ];

  const filteredModels = rawModels.filter((model) => {
    const matchesSearch = model.name.toLowerCase().includes(search.toLowerCase());
    const matchesCategory = selectedCategory === 'All' || model.category === selectedCategory;
    return matchesSearch && matchesCategory;
  });

  return (
    <main>
      <section className={`${styles.hero} page-section`}>
        <div className="container">
          <h2>AI Model Marketplace</h2>
          <p className="small-muted">Discover, evaluate, and purchase trusted AI models with verified performance metrics and community reviews.</p>
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
        <h3>Available Models ({filteredModels.length})</h3>
        <div className={`${styles.grid} grid grid-4`}>
          {filteredModels.map((model, index) => (
            <ModelCard key={index} {...model} />
          ))}
        </div>

        <div style={{ textAlign: 'center', marginTop: 28 }}>
          <div className={styles.pager}>
            <button className={styles.pageActive}>1</button>
            <button>2</button>
            <button>3</button>
            <button>&rsaquo;</button>
          </div>
        </div>
      </section>
    </main>
  );
}