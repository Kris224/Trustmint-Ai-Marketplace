package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client represents an Ethereum blockchain client
type Client struct {
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
	address    common.Address
	chainID    *big.Int
}

// NewClient creates a new blockchain client
func NewClient(rpcURL string, privateKeyHex string) (*Client, error) {
	// Connect to Ethereum node
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %w", err)
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	// Get public key and address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	return &Client{
		client:     client,
		privateKey: privateKey,
		address:    address,
		chainID:    chainID,
	}, nil
}

// GetTransactor creates a transactor for signing transactions
func (c *Client) GetTransactor() (*bind.TransactOpts, error) {
	nonce, err := c.client.PendingNonceAt(context.Background(), c.address)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(500000) // gas limit
	auth.GasPrice = gasPrice

	return auth, nil
}

// GetClient returns the underlying ethclient
func (c *Client) GetClient() *ethclient.Client {
	return c.client
}

// GetAddress returns the wallet address
func (c *Client) GetAddress() common.Address {
	return c.address
}

// Close closes the client connection
func (c *Client) Close() {
	c.client.Close()
}
