
package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "Runs the training process and generates cryptographic hashes.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n--- Starting Secure Training Process ---")

		configHash := os.Getenv("CONFIG_HASH")
		if configHash == "" {
			fmt.Println("\n Error: CONFIG_HASH not found.")
			os.Exit(1)
		}
		fmt.Printf("\n Verified Config Hash: %s\n", configHash)

		viper.SetConfigFile("trustmint.yml")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("\n Error reading config file: %s\n", err)
			os.Exit(1)
		}
		datasetDir := viper.GetString("dataset_dir")
		modelDir := viper.GetString("model_output_dir")
		trainingScript := viper.GetString("training_script")

		fmt.Println("\n- Hashing dataset...")
		datasetHash, err := hashDirectory(datasetDir)
		if err != nil {
			fmt.Printf("\n Error hashing dataset: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("\n- Dataset Hash: %s\n", datasetHash)

		fmt.Println("\n- Running model training script...")
		trainingCmd := exec.Command("sh", "-c", trainingScript)
		trainingCmd.Stdout = os.Stdout
		trainingCmd.Stderr = os.Stderr
		if err := trainingCmd.Run(); err != nil {
			fmt.Printf("\n Error during model training: %s\n", err)
			os.Exit(1)
		}
		fmt.Println("  \n- Training complete.")

		fmt.Println("\n- Hashing output model...")
		modelHash, err := hashDirectory(modelDir)
		if err != nil {
			fmt.Printf("\n Error hashing model: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("  - Model Hash: %s\n", modelHash)

		viper.Set("config_hash", configHash)
		viper.Set("dataset_hash", datasetHash)
		viper.Set("model_hash", modelHash)
		if err := viper.WriteConfigAs(".hashes.yaml"); err != nil {
			fmt.Printf("\n Error saving hashes: %s\n", err)
			os.Exit(1)
		}
		fmt.Println("\n\n  - Proofs saved to .hashes.yaml")

		fmt.Println("\n\n--- ✅ Secure Training Process Finished ---")
	},
}

// hashDirectory function remains the same...
func hashDirectory(path string) (string, error) {
	hasher := sha256.New()
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil { return err }
		if !info.Mode().IsRegular() { return nil }
		file, err := os.Open(filePath)
		if err != nil { return err }
		defer file.Close()
		if _, err := io.Copy(hasher, file); err != nil { return err }
		return nil
	})
	if err != nil { return "", err }
	return hex.EncodeToString(hasher.Sum(nil)), nil
}


func init() {
	rootCmd.AddCommand(trainCmd)
}