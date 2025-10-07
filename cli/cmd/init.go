package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// variables to store flag values
var (
	walletAddr string
	rpcEndpt   string
	storage    string
	outFile    string
)

func init() {
	// attach initCmd to root
	rootCmd.AddCommand(initCmd)

	// add flags
	initCmd.Flags().StringVar(&walletAddr, "wallet", "", "wallet address (required)")
	initCmd.Flags().StringVar(&rpcEndpt, "rpc", "", "blockchain RPC endpoint (required)")
	initCmd.Flags().StringVar(&storage, "storage", "ipfs", "storage preference (ipfs|local)")
	initCmd.Flags().StringVar(&outFile, "out", "trustmint.yaml", "output config yaml file")

	// make wallet and rpc required
	initCmd.MarkFlagRequired("wallet")
	initCmd.MarkFlagRequired("rpc")
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize TrustMint config",
	RunE: func(cmd *cobra.Command, args []string) error {
		// very basic YAML output (manual write, not Viper yet)
		content := fmt.Sprintf("wallet_address: %s\nrpc_endpoint: %s\nstorage: %s\n",
			walletAddr, rpcEndpt, storage)

		if err := os.WriteFile(outFile, []byte(content), 0644); err != nil {
			return err
		}
		fmt.Printf("Config saved to %s ✅\n", outFile)
		return nil
	},
}
