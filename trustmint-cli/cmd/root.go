// cmd/root.go
package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "trustmint",
    Short: "A CLI to securely train and publish AI models with blockchain provenance.",
    Long: `Trustmint is a tool that ensures AI models are trained in a secure,
isolated environment, with their lineage tracked on the Polygon blockchain.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Welcome to Trustmint CLI! You are inside the secure container.")
        fmt.Println("Run 'trustmint --help' to see available commands.")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}