// src/pages/Profile.jsx
import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { auth } from '../firebase';
import { onAuthStateChanged } from 'firebase/auth';
import styles from './Login.module.css';  // Reuse for consistent card design

export default function Profile() {
  const [user, setUser] = useState(null);
  const [displayName, setDisplayName] = useState('');
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, (currentUser) => {
      if (!currentUser) {
        navigate('/login');
      } else {
        setUser(currentUser);
        setDisplayName(currentUser.displayName || currentUser.email.split('@')[0]);
        setLoading(false);
      }
    });
    return () => unsubscribe();
  }, [navigate]);

  const handleUpdateName = async () => {
    if (!user || !displayName) return;
    try {
      await user.updateProfile({ displayName });
      alert('Profile updated!');
    } catch (error) {
      alert('Update failed: ' + error.message);
    }
  };

  if (loading) return <div className={styles.container}>Loading...</div>;

  return (
    <main className={styles.container}>
      <div className={styles.card}>
        <h2 className={styles.title}>Your Profile</h2>
        <p className={styles.subtitle}>Manage your account details</p>
        <div className={styles.inputGroup}>
          <label>Email</label>
          <input type="email" value={user.email} disabled className="input" />
        </div>
        <div className={styles.inputGroup}>
          <label>Display Name</label>
          <input
            type="text"
            value={displayName}
            onChange={(e) => setDisplayName(e.target.value)}
            placeholder="Enter display name"
          />
        </div>
        <button onClick={handleUpdateName} className={styles.loginBtn}>
          Update Profile
        </button>
        <p className={styles.footerText}>
          Joined on: {user.metadata.creationTime?.split('T')[0] || 'Unknown'}
        </p>
      </div>
    </main>
  );
}