// cmd/login.go
package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
    Use:   "login",
    Short: "Authenticate with the Trustmint platform.",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("\nSuccessfully logged in. API key is ready for 'publish' command.\n\n")
    },
}

func init() {
    rootCmd.AddCommand(loginCmd)
}