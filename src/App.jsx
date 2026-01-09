import React, { useEffect, useState } from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';  // No <Router> here
import Home from './pages/Home';
import Marketplace from './pages/Marketplace';
import Developer from './pages/Developer';
import Login from './pages/Login';
import Signup from './pages/Signup';
import Navbar from './components/Navbar';
import Footer from './components/Footer';
import { auth } from './firebase';
import { onAuthStateChanged } from 'firebase/auth';

export default function App() {
  const [theme, setTheme] = useState(() => {
    const saved = localStorage.getItem('theme');
    return saved || (window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light');
  });
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    document.documentElement.setAttribute('data-theme', theme);
    localStorage.setItem('theme', theme);
  }, [theme]);

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, (user) => {
      setIsLoggedIn(!!user);
    });
    return () => unsubscribe();
  }, []);

  // Placeholder for Profile (create src/pages/Profile.jsx if missing)
  const Profile = () => <div style={{ padding: '2rem', textAlign: 'center' }}><h1>Profile</h1><p>Coming soon!</p></div>;

  return (
    <>
      <Navbar theme={theme} setTheme={setTheme} isLoggedIn={isLoggedIn} />
      <div style={{ minHeight: 'calc(100vh - 140px)' }}>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={!isLoggedIn ? <Login onLogin={() => setIsLoggedIn(true)} /> : <Navigate to="/developer" />} />
          <Route path="/signup" element={!isLoggedIn ? <Signup /> : <Navigate to="/developer" />} />
          <Route path="/marketplace" element={isLoggedIn ? <Marketplace /> : <Navigate to="/login" />} />
          <Route path="/developer" element={isLoggedIn ? <Developer /> : <Navigate to="/login" />} />
          <Route path="/profile" element={isLoggedIn ? <Profile /> : <Navigate to="/login" />} />
        </Routes>
      </div>
      <Footer />
    </>
  );
}