import React from "react";
import styles from "./Login.module.css";
import { useNavigate } from "react-router-dom";

export default function Login() {
  const navigate = useNavigate();

  function handleLogin(e) {
    e.preventDefault();
    // Placeholder: in a real app you'd send credentials to backend
    alert("Logged in (placeholder).");
    navigate("/developer");
  }

  function handleGoogleLogin() {
    // Placeholder for Google Sign-In
    alert(
      "Google Login clicked — this is a placeholder. Add your Google OAuth client integration."
    );
    navigate("/developer");
  }

  return (
    <main className={styles.container}>
      <div className={styles.card}>
        <h2 className={styles.title}>Welcome Back</h2>
        <p className={styles.subtitle}>Sign in to your Trustmint account</p>

        <form onSubmit={handleLogin} className={styles.form}>
          <div className={styles.inputGroup}>
            <label>Email</label>
            <input type="email" required placeholder="you@example.com" />
          </div>

          <div className={styles.inputGroup}>
            <label>Password</label>
            <input type="password" required placeholder="Your password" />
          </div>

          <button type="submit" className={styles.loginBtn}>
            Login
          </button>
        </form>

        <div className={styles.divider}>
          <span>or</span>
        </div>

        <button
          type="button"
          className={styles.googleBtn}
          onClick={handleGoogleLogin}
        >
          <img src="/google-icon.svg" alt="Google" />
          Login with Google
        </button>

        <p className={styles.footerText}>
          Don’t have an account? <a href="#">Sign up</a>
        </p>
      </div>
    </main>
  );
}
