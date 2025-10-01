import React from 'react'
import { Link, useLocation } from 'react-router-dom'
import styles from './Navbar.module.css'

export default function Navbar({theme,setTheme}){
  const location = useLocation()
  return (
    <header className={styles.header}>
      <div className={styles.inner}>
        <div className={styles.brand}>
          <Link to="/">Trustmint</Link>
        </div>
        <nav className={styles.nav}>
          <Link className={location.pathname==='/'?styles.active:''} to="/">Home</Link>
          <Link className={location.pathname==='/marketplace'?styles.active:''} to="/marketplace">Marketplace</Link>
          <Link className={location.pathname==='/developer'?styles.active:''} to="/developer">Developer</Link>
        </nav>
        <div className={styles.actions}>
          <button className={styles.themeBtn} onClick={()=>setTheme(theme==='dark'?'light':'dark')}>
            {theme==='dark' ? '☀️' : '🌙'}
          </button>
          <Link to="/login" className={styles.loginBtn}>Login</Link>
        </div>
      </div>
    </header>
  )
}
