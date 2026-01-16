package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Launches the secure training environment using Docker.",
	Long:  `Parses trustmint.yml and runs a secure Docker container with the specified configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("▶️  Starting Trustmint Workflow...")

		// 1. Check for Config
		configFile := "trustmint.yml"
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			fmt.Println("❌ Error: 'trustmint.yml' not found. Run 'trustmint init' first.")
			os.Exit(1)
		}

		// 2. Check Docker
		if _, err := exec.LookPath("docker"); err != nil {
			fmt.Println("❌ Error: Docker is not installed or not in PATH.")
			os.Exit(1)
		}

		// 3. Parse Config
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("❌ Error reading config file: %v\n", err)
			os.Exit(1)
		}

		networkMode := viper.GetString("environment.network")
		readOnlyRoot := viper.GetBool("environment.read_only_root")

		var volumes []struct {
			Source   string `mapstructure:"source"`
			Target   string `mapstructure:"target"`
			ReadOnly bool   `mapstructure:"read_only"`
		}
		if err := viper.UnmarshalKey("environment.volumes", &volumes); err != nil {
			fmt.Printf("❌ Error parsing volumes: %v\n", err)
			os.Exit(1)
		}

		// 4. Construct Docker Command
		dockerImage := "20demo/trustmint-cli:latest" // Should make this configurable or versioned?

		dockerArgs := []string{"run", "-it", "--rm"}

		// Network
		if networkMode != "" {
			dockerArgs = append(dockerArgs, "--network", networkMode)
		} else {
			// Default to none if not specified for security
			dockerArgs = append(dockerArgs, "--network", "none")
		}

		// Read Only Root
		if readOnlyRoot {
			dockerArgs = append(dockerArgs, "--read-only")
		}

		// Volumes
		cwd, _ := os.Getwd()
		for _, vol := range volumes {
			absSource := filepath.Join(cwd, vol.Source)
			mount := fmt.Sprintf("%s:%s", absSource, vol.Target)
			if vol.ReadOnly {
				mount += ":ro"
			}
			dockerArgs = append(dockerArgs, "-v", mount)
		}

		// Pass Config Hash (Just for backward compat/display, mostly verified internally now)
		// We calculate it here just to be consistent, but the real check is inside.
		// Actually, let's skip passing CONFIG_HASH env var since the new security model
		// relies on internal hashing. But to keep the container happy if it expects it:
		// (Wait, the *old* container script expected it. The new container logic will handle it.)
		// We will pass it as a dummy or calculated value just in case.
		// Actually, let's keep it clean. No CONFIG_HASH env var.

		dockerArgs = append(dockerArgs, dockerImage)

		fmt.Println("⚙️  Configuration loaded.")
		fmt.Printf("🚀 Launching container (%s)...\n", dockerImage)
		fmt.Println("----------------------------------------------------")

		// 5. Execute
		c := exec.Command("docker", dockerArgs...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		if err := c.Run(); err != nil {
			fmt.Printf("\n❌ Error running container: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("\n✅ Secure session ended.")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
