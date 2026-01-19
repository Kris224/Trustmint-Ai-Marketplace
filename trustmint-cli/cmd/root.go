// cmd/root.go
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "trustmint",
	Short: "A CLI to securely initialize, train, and publish AI models.",
	Long: `Trustmint is a tool that ensures AI models are trained in a secure,
isolated environment, with their lineage tracked on the blockchain.

Available Commands:
  init      Initialize a new project (Host)
  start     Launch the secure training environment (Host)
  train     Run the training script and sign artifacts (Container)
  publish   Upload artifacts to the marketplace (Host)`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
