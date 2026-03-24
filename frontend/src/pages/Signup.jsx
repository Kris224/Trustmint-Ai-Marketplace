// src/pages/Signup.jsx
import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import styles from './Login.module.css';
import googleLogo from '../assets/download.png';
import { auth } from '../firebase';
import {
    createUserWithEmailAndPassword,
    GoogleAuthProvider,
    signInWithPopup
} from 'firebase/auth';

export default function Signup() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
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

    const handleSignup = async (e) => {
        e.preventDefault();

        if (!email || !password || !confirmPassword) {
            return showToast('Please fill all fields', 'error');
        }

        if (password !== confirmPassword) {
            return showToast('Passwords do not match!', 'error');
        }

        if (password.length < 6) {
            return showToast('Password must be at least 6 characters', 'error');
        }

        setLoading(true);
        try {
            await createUserWithEmailAndPassword(auth, email, password);
            showToast('Account created successfully! 🎉');
            navigate('/developer');
        } catch (error) {
            let errorMsg = 'Signup failed';
            if (error.code === 'auth/email-already-in-use') {
                errorMsg = 'Email already in use';
            } else if (error.code === 'auth/weak-password') {
                errorMsg = 'Password is too weak';
            } else if (error.code === 'auth/invalid-email') {
                errorMsg = 'Invalid email address';
            }
            showToast(errorMsg, 'error');
        } finally {
            setLoading(false);
        }
    };

    const handleGoogleSignup = async () => {
        const googleProvider = new GoogleAuthProvider();
        setLoading(true);
        try {
            await signInWithPopup(auth, googleProvider);
            showToast('Signed up with Google successfully! 🎉');
            navigate('/developer');
        } catch (error) {
            showToast('Google signup failed: ' + error.message, 'error');
            setLoading(false);
        }
    };

    return (
        <main className={styles.container}>
            <div className={styles.card}>
                <h2 className={styles.title}>Create Account</h2>
                <p className={styles.subtitle}>Join Trustmint to start trading AI models</p>

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
                            placeholder="At least 6 characters"
                        />
                    </div>

                    <div className={styles.inputGroup}>
                        <label>Confirm Password</label>
                        <input
                            type="password"
                            value={confirmPassword}
                            onChange={(e) => setConfirmPassword(e.target.value)}
                            required
                            placeholder="Re-enter your password"
                        />
                    </div>

                    <button type="submit" className={styles.loginBtn} disabled={loading}>
                        {loading ? <div className="spinner"></div> : 'Create Account'}
                    </button>
                </form>

                <div className={styles.divider}>
                    <span>or</span>
                </div>

                <div style={{ display: 'flex', justifyContent: 'center' }}>
                    <button
                        type="button"
                        className={styles.googleBtn}
                        onClick={handleGoogleSignup}
                        disabled={loading}
                    >
                        <img src={googleLogo} alt="Google" />
                        Sign up with Google
                    </button>
                </div>

                <p className={styles.footerText}>
                    Already have an account? <Link to="/login">Sign in</Link>
                </p>
            </div>
        </main>
    );
}
