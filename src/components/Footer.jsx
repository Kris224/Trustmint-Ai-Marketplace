import React from 'react'
import styles from './Footer.module.css'

export default function Footer(){
  return (
    <footer className={styles.footer}>
      <div className="container">
        <div>
          <h3>Trustmint</h3>
          <p>The leading decentralized AI model marketplace for developers and businesses worldwide.</p>
        </div>
        <div>
          <h4>Company</h4>
          <ul>
            <li>About</li>
            <li>Contact</li>
          </ul>
        </div>
        <div>
          <h4>Resources</h4>
          <ul>
            <li>Docs</li>
            <li>GitHub</li>
          </ul>
        </div>
      </div>
      <div className="credit">© 2025 Trustmint. All rights reserved.</div>
    </footer>
  )
}
