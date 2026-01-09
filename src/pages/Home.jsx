// src/pages/Home.jsx
import React from 'react';
import { Link } from 'react-router-dom';
import styles from './Home.module.css';

export default function Home() {
  return (
    <main>
      <section className={`${styles.hero} page-section`}>
        <div className="container">
          <h1>Trustmint – Decentralized AI Model Marketplace</h1>
          <p className={styles.lead}>Discover, deploy, and monetize AI models with confidence. Our platform ensures quality, transparency, and seamless integration.</p>
          <div className={styles.cta}>
            <Link to="/marketplace" className={styles.primary}>Get Started</Link>
          </div>
        </div>
      </section>

      <section className="page-section">
        <div className="container">
          <div className={`${styles.stats} grid grid-4`}>
            <div className="card">
              <h2>247</h2>
              <p className="small-muted">Available Models</p>
            </div>
            <div className="card">
              <h2>12k+</h2>
              <p className="small-muted">Happy Customers</p>
            </div>
            <div className="card">
              <h2>4.9</h2>
              <p className="small-muted">Avg Rating</p>
            </div>
            <div className="card">
              <h2>99.9%</h2>
              <p className="small-muted">Uptime</p>
            </div>
          </div>
        </div>
      </section>

      <section className="page-section">
        <div className="container">
          <h3 style={{ textAlign: "center", marginBottom: "1rem" }}>Why Choose Trustmint?</h3>
          <p className="small-muted" style={{ textAlign: "center", marginBottom: "2rem" }}>
            Experience the future of AI model marketplace with our trusted platform
          </p>
          <div className={`${styles.features} grid grid-3`}>
            <div className="card">
              <h4>Verified Models</h4>
              <p className="small-muted">Every AI model undergoes rigorous quality checks and verification processes.</p>
            </div>
            <div className="card">
              <h4>Trust Scores</h4>
              <p className="small-muted">Transparent performance metrics and community ratings help you make informed decisions.</p>
            </div>
            <div className="card">
              <h4>Easy Integration</h4>
              <p className="small-muted">Seamless integration with comprehensive APIs and SDKs.</p>
            </div>
          </div>
        </div>
      </section>
    </main>
  );
}