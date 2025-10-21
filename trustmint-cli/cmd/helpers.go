// cmd/helpers.go
package cmd

import (
	"fmt"
	"strings"
)

func printHeader(title string) {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("== %s\n", strings.ToUpper(title))
	fmt.Println(strings.Repeat("=", 60))
}

func printSuccess(message string) {
	fmt.Printf("\n🎉 SUCCESS: %s\n", message)
}

func printError(message string) {
	fmt.Printf("\n❌ ERROR: %s\n", message)
}