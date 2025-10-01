import React from 'react'
import styles from './Marketplace.module.css'

const ModelCard = ({name,price,score,category}) => (
  <div className={styles.card}>
    <div className={styles.icon}></div>
    <div>
      <h4>{name}</h4>
      <p className="small-muted">Trust Score: {score} <span style={{float:'right'}}>★ 4.7</span></p>
      <div style={{marginTop:12}}>
        <strong style={{fontSize:20}}>${price}</strong>
        <button className={styles.viewBtn}>View Details</button>
      </div>
    </div>
  </div>
)

export default function Marketplace(){
  return (
    <main>
      <section className={styles.hero}>
        <div className="container">
          <h2>AI Model Marketplace</h2>
          <p className="small-muted">Discover, evaluate, and purchase trusted AI models with verified performance metrics and community reviews.</p>
          <div className={styles.searchRow}>
            <input placeholder="Search AI models..." />
            <select><option>All Categories</option></select>
            <button className="link-btn">Search</button>
          </div>
        </div>
      </section>

      <section className="container" style={{paddingTop:24}}>
        <div className={styles.filters}>
          <button className={styles.filterActive}>All</button>
          <button>NLP</button>
          <button>Vision</button>
          <button>Speech</button>
          <button>Data Processing</button>
        </div>
      </section>

      <section className="container" style={{paddingTop:24}}>
        <h3>Available Models</h3>
        <div className={styles.grid}>
          <ModelCard name="TextGenX" price="49" score="92" />
          <ModelCard name="VisionPro" price="39" score="88" />
          <ModelCard name="SpeechAI Lite" price="59" score="95" />
          <ModelCard name="DataSummarizer" price="29" score="85" />
        </div>

        <div style={{textAlign:'center', marginTop:28}}>
          <div className={styles.pager}>
            <button className={styles.pageActive}>1</button>
            <button>2</button>
            <button>3</button>
          </div>
        </div>
      </section>
    </main>
  )
}
