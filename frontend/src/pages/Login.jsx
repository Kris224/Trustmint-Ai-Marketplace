import React from 'react'
import styles from './Login.module.css'
import { useNavigate } from 'react-router-dom'

export default function Login(){
  const navigate = useNavigate()

  function handleLogin(e){
    e.preventDefault()
    // Placeholder: in a real app you'd send credentials to backend
    alert('Logged in (placeholder).')
    navigate('/developer')
  }

  function handleGoogleLogin(){
    // Placeholder for Google Sign-In
    // In production: integrate Google Identity Services or OAuth2 flow.
    alert('Google Login clicked — this is a placeholder. Add your Google OAuth client integration.');
    navigate('/developer')
  }

  return (
    <main>
      <section className="container" style={{maxWidth:720}}>
        <div className="card" style={{padding:28}}>
          <h2 style={{marginTop:0}}>Login</h2>
          <p className="small-muted">Sign in to your Trustmint account</p>

          <form onSubmit={handleLogin} className={styles.form}>
            <label>Email
              <input type="email" required placeholder="you@example.com" />
            </label>
            <label>Password
              <input type="password" required placeholder="Your password" />
            </label>
            <div className={styles.row}>
              <button type="submit" className="primary">Login</button>
              <button type="button" className="link-btn" onClick={handleGoogleLogin}>Login with Google</button>
            </div>
          </form>
          <p className="small-muted" style={{marginTop:12}}>Note: Google login button is a client-side placeholder. To enable real Google OAuth, follow the README steps to configure Google Identity Services or OAuth on your backend.</p>
        </div>
      </section>
    </main>
  )
}
