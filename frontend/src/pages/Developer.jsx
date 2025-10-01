import React from 'react'
import styles from './Developer.module.css'

export default function Developer(){
  return (
    <main>
      <section className="container">
        <h2>Developer Dashboard</h2>
        <button className="link-btn" style={{marginTop:8}}>Download CLI</button>

        <h3 style={{marginTop:28}}>Quick Start Guide</h3>
        <div className={styles.steps}>
          <div className="card">
            <h4>1. Download CLI</h4>
            <p className="small-muted">Install the Trustmint CLI tool to manage your models</p>
          </div>
          <div className="card">
            <h4>2. Authenticate</h4>
            <p className="small-muted">Connect your account using API credentials</p>
          </div>
          <div className="card">
            <h4>3. Upload Model</h4>
            <p className="small-muted">Submit your model with complete metadata</p>
          </div>
          <div className="card">
            <h4>4. Wait for Verification</h4>
            <p className="small-muted">Our team reviews your model (24-48 hours)</p>
          </div>
          <div className="card">
            <h4>5. Start Earning</h4>
            <p className="small-muted">Your verified model goes live and starts generating revenue</p>
          </div>
        </div>

        <h3 style={{marginTop:28}}>Earnings Overview</h3>
        <div style={{display:'grid',gridTemplateColumns:'repeat(4,1fr)', gap:16, marginTop:12}}>
          <div className="card"><p className="small-muted">This Month</p><h3>$2,450.00</h3></div>
          <div className="card"><p className="small-muted">Last Month</p><h3>$1,890.00</h3></div>
          <div className="card"><p className="small-muted">Total Earned</p><h3>$12,340.00</h3></div>
          <div className="card"><p className="small-muted">Pending Payout</p><h3>$340.00</h3></div>
        </div>

        <h3 style={{marginTop:28}}>Uploaded Models</h3>
        <div className="card" style={{marginTop:12}}>
          <table style={{width:'100%', borderCollapse:'collapse'}}>
            <thead style={{textAlign:'left'}}>
              <tr><th>Model Name</th><th>Accuracy</th><th>Status</th><th>Actions</th></tr>
            </thead>
            <tbody>
              <tr style={{borderTop:'1px solid rgba(0,0,0,0.06)'}}><td>ImageNet Classifier</td><td>92%</td><td>Verified</td><td>⋮</td></tr>
              <tr style={{borderTop:'1px solid rgba(0,0,0,0.06)'}}><td>Text Summarizer</td><td>88%</td><td>Pending</td><td>⋮</td></tr>
              <tr style={{borderTop:'1px solid rgba(0,0,0,0.06)'}}><td>Speech2Text</td><td>95%</td><td>Verified</td><td>⋮</td></tr>
            </tbody>
          </table>
        </div>
      </section>
    </main>
  )
}
