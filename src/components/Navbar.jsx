// src/components/Navbar.jsx
import React, { useState, useEffect, useRef } from 'react';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import { auth } from '../firebase';
import ConnectWallet from './ConnectWallet';  // Existing component
import styles from './Navbar.module.css';

export default function Navbar({ theme, setTheme, isLoggedIn }) {
  const [userOpen, setUserOpen] = useState(false);
  const location = useLocation();
  const navigate = useNavigate();
  const dropdownRef = useRef(null);

  const handleLogout = async () => {
    await auth.signOut();
    setUserOpen(false);
    navigate('/');
  };

  useEffect(() => {
    const handleClickOutside = (event) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target)) {
        setUserOpen(false);
      }
    };
    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, []);

  return (
    <header className={styles.header}>
      <div className="container">
        <div className={styles.inner}>
          <Link to="/" className={styles.brand}>Trustmint</Link>
          <nav className={styles.nav}>
            <Link className={location.pathname === '/' ? styles.active : ''} to="/">Home</Link>
            <Link className={location.pathname === '/marketplace' ? styles.active : ''} to="/marketplace">Marketplace</Link>
            <Link className={location.pathname === '/developer' ? styles.active : ''} to="/developer">Developer</Link>
          </nav>
          <div className={styles.actions}>
            <button className={styles.themeBtn} onClick={() => setTheme(theme === 'dark' ? 'light' : 'dark')} title="Toggle theme">
              {theme === 'dark' ? '☀️' : '🌙'}
            </button>
            <ConnectWallet isLoggedIn={isLoggedIn} />
            {isLoggedIn ? (
              <div className={styles.userMenu} ref={dropdownRef}>
                <button className={styles.userBtn} onClick={() => setUserOpen(!userOpen)} title="User menu">
                  👤
                </button>
                {userOpen && (
                  <div className={styles.dropdown}>
                    <Link to="/profile" className={styles.dropdownItem} onClick={() => setUserOpen(false)}>Profile</Link>
                    <button onClick={handleLogout} className={styles.dropdownItem}>Logout</button>
                  </div>
                )}
              </div>
            ) : (
              <Link to="/login" className={styles.loginBtn}>Login</Link>
            )}
          </div>
        </div>
      </div>
    </header>
  );
}