import { useState, useEffect } from "react";
import { ethers } from "ethers";

function ConnectWallet() {
  const [account, setAccount] = useState(null);

  useEffect(() => {
    // Auto-detect wallet if already connected
    if (window.ethereum) {
      window.ethereum.request({ method: "eth_accounts" }).then((accounts) => {
        if (accounts.length > 0) setAccount(accounts[0]);
      });

      window.ethereum.on("accountsChanged", (accounts) => {
        setAccount(accounts[0] || null);
      });
    }
  }, []);

  async function connectWallet() {
    if (!window.ethereum) {
      alert("MetaMask not detected. Please install it.");
      return;
    }

    try {
      const provider = new ethers.BrowserProvider(window.ethereum);
      const accounts = await provider.send("eth_requestAccounts", []);
      setAccount(accounts[0]);
      console.log("Connected wallet:", accounts[0]);
    } catch (err) {
      console.error("MetaMask connection error:", err);
    }
  }

  return (
    <div style={{ padding: "10px" }}>
      {account ? (
        <p>✅ Connected Wallet: {account}</p>
      ) : (
        <button
          onClick={connectWallet}
          style={{
            backgroundColor: "#6366f1",
            color: "white",
            padding: "8px 16px",
            border: "none",
            borderRadius: "6px",
            cursor: "pointer",
          }}
        >
          Connect MetaMask
        </button>
      )}
    </div>
  );
}

export default ConnectWallet;
