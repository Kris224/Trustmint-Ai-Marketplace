import React from 'react';
import { Link } from 'react-router-dom';
import styles from './Developer.module.css';

export default function Developer() {
  return (
    <main>
      <section className="page-section">
        <div className="container">
          <div className={styles.dashboardHeader}>
            <h2>Developer Dashboard</h2>
            <Link to="/cli-download" className="btn btn-primary">Download CLI</Link>
          </div>

          <h3>Quick Start Guide</h3>
          <div className={`${styles.steps} grid grid-5`}>
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

          <h3>Uploaded Models</h3>
          <div className="card">
            <div className={styles.tableHeader}>
              <Link to="/upload" className="btn btn-secondary">Upload New</Link>
            </div>
            <table className="table">
              <thead>
                <tr>
                  <th>Model Name</th>
                  <th>Accuracy</th>
                  <th>Status</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>ImageNet Classifier</td>
                  <td>92%</td>
                  <td><span className={styles.statusVerified}>Verified</span></td>
                  <td><button className="btn btn-primary" style={{ padding: '0.25rem 0.5rem', fontSize: '0.75rem' }}>⋮</button></td>
                </tr>
                <tr>
                  <td>Text Summarizer</td>
                  <td>88%</td>
                  <td><span className={styles.statusPending}>Pending</span></td>
                  <td><button className="btn btn-primary" style={{ padding: '0.25rem 0.5rem', fontSize: '0.75rem' }}>⋮</button></td>
                </tr>
                <tr>
                  <td>Speech2Text</td>
                  <td>95%</td>
                  <td><span className={styles.statusVerified}>Verified</span></td>
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