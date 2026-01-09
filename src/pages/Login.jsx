// src/pages/Login.jsx
import React, { useState, useEffect } from "react";
import styles from "./Login.module.css";
import { useNavigate, Link } from "react-router-dom";
import googleLogo from '../assets/download.png';
import { auth } from '../firebase';
import { 
  signInWithEmailAndPassword, 
  sendPasswordResetEmail, 
  GoogleAuthProvider, 
  signInWithPopup, 
  sendSignInLinkToEmail,
  isSignInWithEmailLink,
  signInWithEmailLink
} from 'firebase/auth';

export default function Login({ onLogin }) {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const showToast = (message, type = 'success') => {
    const toast = document.createElement('div');
    toast.className = `toast ${type}`;
    toast.textContent = message;
    document.body.appendChild(toast);
    setTimeout(() => toast.classList.add('show'), 100);
    setTimeout(() => {
      toast.classList.remove('show');
      setTimeout(() => toast.remove(), 300);
    }, 3000);
  };

  const handleLogin = async (e) => {
    e.preventDefault();
    if (!email || !password) return showToast('Enter email and password', 'error');
    setLoading(true);
    try {
      await signInWithEmailAndPassword(auth, email, password);
      showToast('Logged in successfully!');
      if (onLogin) onLogin();
      navigate("/developer");
    } catch (error) {
      showToast('Login failed: ' + error.message, 'error');
    } finally {
      setLoading(false);
    }
  };

  const handleForgotPassword = async (e) => {
    e.preventDefault();
    if (!email) return showToast('Enter your email first!', 'error');
    try {
      await sendPasswordResetEmail(auth, email);
      showToast('Reset link sent to your email!');
    } catch (error) {
      showToast('Error: ' + error.message, 'error');
    }
  };

  const handleGoogleLogin = async () => {
    const googleProvider = new GoogleAuthProvider();
    setLoading(true);
    try {
      await signInWithPopup(auth, googleProvider);
      showToast('Google login successful!');
      if (onLogin) onLogin();
      navigate("/developer");
    } catch (error) {
      showToast('Google login failed: ' + error.message, 'error');
      setLoading(false);
    }
  };

  const handleMagicLink = async (e) => {
    e.preventDefault();
    if (!email) return showToast('Enter your email!', 'error');
    const actionCodeSettings = {
      url: window.location.href,
      handleCodeInApp: true
    };
    try {
      await sendSignInLinkToEmail(auth, email, actionCodeSettings);
      window.localStorage.setItem('emailForSignIn', email);
      showToast('Magic link sent! Check your email.');
    } catch (error) {
      showToast('Error: ' + error.message, 'error');
    }
  };

  useEffect(() => {
    if (isSignInWithEmailLink(auth, window.location.href)) {
      let storedEmail = window.localStorage.getItem('emailForSignIn');
      if (!storedEmail) storedEmail = window.prompt('Enter your email for confirmation');
      if (storedEmail) {
        signInWithEmailLink(auth, storedEmail, window.location.href)
          .then(() => {
            window.localStorage.removeItem('emailForSignIn');
            showToast('Logged in via magic link!');
            navigate("/developer");
          })
          .catch((error) => showToast('Error: ' + error.message, 'error'));
      }
    }
  }, [navigate]);

  return (
    <main className={styles.container}>
      <div className={styles.card}>
        <h2 className={styles.title}>Welcome Back</h2>
        <p className={styles.subtitle}>Sign in to your Trustmint account</p>

        <form onSubmit={handleLogin} className={styles.form}>
          <div className={styles.inputGroup}>
            <label>Email</label>
            <input 
              type="email" 
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required 
              placeholder="you@example.com" 
            />
          </div>

          <div className={styles.inputGroup}>
            <label>Password</label>
            <input 
              type="password" 
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required 
              placeholder="Your password" 
            />
          </div>

          <button type="submit" className={styles.loginBtn} disabled={loading}>
            {loading ? <div className="spinner"></div> : 'Login'}
          </button>

          <button type="button" onClick={handleForgotPassword} className={styles.googleBtn} style={{ background: 'white', color: '#111827', marginTop: '10px' }}>
            Forgot Password?
          </button>
          <button type="button" onClick={handleMagicLink} className={styles.googleBtn} style={{ background: 'linear-gradient(135deg, var(--primary), var(--secondary))', color: 'white', marginTop: '5px', boxShadow: '0 4px 12px rgba(99, 102, 241, 0.3)' }}>
            Send Magic Link
          </button>
        </form>

        <div className={styles.divider}>
          <span>or</span>
        </div>
        <div style={{ display: 'flex', justifyContent: 'center' }}>
          <button
            type="button"
            className={styles.googleBtn}
            onClick={handleGoogleLogin}
            disabled={loading}
          >
            <img src={googleLogo} alt="Google" />
            Login with Google
          </button>
        </div>

        <p className={styles.footerText}>
          Don’t have an account? <Link to="/signup">Sign up</Link>
        </p>
      </div>
    </main>
  );
}