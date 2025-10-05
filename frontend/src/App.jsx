import React, { useEffect, useState } from 'react'
import { Routes, Route, Link } from 'react-router-dom'
import Home from './pages/Home'
import Marketplace from './pages/Marketplace'
import Developer from './pages/Developer'
import Login from './pages/Login'
import Navbar from './components/Navbar'
import Footer from './components/Footer'
import ConnectWallet from "./components/ConnectWallet";

export default function App(){
  const [theme, setTheme] = useState(localStorage.getItem('theme') || 'light')
  useEffect(()=>{
    document.documentElement.setAttribute('data-theme', theme === 'dark' ? 'dark' : 'light')
    localStorage.setItem('theme', theme)
  },[theme])

  return (
    <>
      <Navbar theme={theme} setTheme={setTheme} />
      <ConnectWallet />
      <div style={{minHeight:'calc(100vh - 220px)'}}>
        <Routes>
          <Route path='/' element={<Home/>} />
          <Route path='/marketplace' element={<Marketplace/>} />
          <Route path='/developer' element={<Developer/>} />
          <Route path='/login' element={<Login/>} />
        </Routes>
      </div>
      <Footer />
    </>
  )
}
