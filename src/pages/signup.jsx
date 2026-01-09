// src/pages/Signup.jsx
import React, { useState } from "react";
import styles from "./Login.module.css";  // Keep reusing
import { useNavigate, Link } from "react-router-dom";  // Add Link for better nav
import { auth } from '../firebase';  // Add this
import { createUserWithEmailAndPassword } from 'firebase/auth';  // Add this

export default function Signup() {
  const [email, setEmail] = useState('');  // Add states
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  // Toast function
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

  function handleSignup(e) {
    e.preventDefault();
    if (!email || !password) return showToast('Enter email and password', 'error');
    setLoading(true);
    createUserWithEmailAndPassword(auth, email, password)
      .then(() => {
        setLoading(false);
        showToast('Account created! Welcome aboard.');
        navigate("/developer");
      })
      .catch((error) => {
        setLoading(false);
        showToast('Sign-up failed: ' + error.message, 'error');
      });
  }

  return (
    <main className={styles.container}>
      <div className={styles.card}>
        <h2 className={styles.title}>Create Account</h2>
        <p className={styles.subtitle}>Sign up for your Trustmint account</p>

        <form onSubmit={handleSignup} className={styles.form}>
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
            {loading ? <div className="spinner"></div> : 'Sign Up'}
          </button>
        </form>

        <p className={styles.footerText}>
          Already have an account? <Link to="/login">Login</Link>  {/* Use Link instead of <a> */}
        </p>
      </div>
    </main>
  );
}