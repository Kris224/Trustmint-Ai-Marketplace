import React from 'react'
import styles from './Home.module.css'

export default function Home(){
  return (
    <main>
      <section className={styles.hero}>
        <div className="container">
          <h1>Trustmint – Decentralized AI Model Marketplace</h1>
          <p className={styles.lead}>Discover, deploy, and monetize AI models with confidence. Our platform ensures quality, transparency, and seamless integration.</p>
          <div className={styles.cta}>
            <button className={styles.primary}>Get Started</button>
            <button className={styles.ghost}>Login</button>
          </div>
        </div>
      </section>

      <section className="container" style={{paddingTop:36}}>
        <div className={styles.stats}>
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
      </section>

      <section className="container" style={{paddingTop:48}}>
        <h3 style={{textAlign:'center'}}>Why Choose Trustmint?</h3>
        <p className="small-muted" style={{textAlign:'center'}}>Experience the future of AI model marketplace with our trusted platform</p>
        <div className={styles.features}>
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
      </section>
    </main>
  )
}
