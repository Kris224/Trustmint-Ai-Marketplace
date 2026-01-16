package cmd

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var AppPrivateKey string // Injected at build time via -ldflags

var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "Runs the training process and generates cryptographic proofs.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n--- Starting Secure Training Process ---")

		fmt.Println("\n- verifying configuration integrity...")
		// 1. Hash trustmint.yml directly (Source of Truth)
		configHash, err := hashFile("trustmint.yml")
		if err != nil {
			fmt.Printf("\n Error hashing config file: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("  - Config Hash: %s\n", configHash)

		viper.SetConfigFile("trustmint.yml")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("\n Error reading config file: %s\n", err)
			os.Exit(1)
		}
		datasetDir := viper.GetString("dataset_dir")
		modelDir := viper.GetString("model_output_dir")
		trainingScript := viper.GetString("training_script")
		// Extract script filename (e.g. "python3 train.py" -> "train.py")
		// This is a naive extraction, assuming the last part is the file.
		// A better way is to enforce "script_file" in config, but for now we parse.

		// Let's just assume the user puts the script file in the project root for now.
		// Actually, let's just hash the 'train.py' file if it exists, or whatever is specified.
		// Simpler: assume the script filename is 'train.py' for this demo or parse it properly.
		// Let's just look for "train.py" since our init creates it.
		// Or better, read the file specified in the command string? No, that's hard to parse reliably.
		// Let's rely on convention: "train.py" for now, or just hash the file 'train.py' explicitly.
		scriptFile := "train.py"

		fmt.Println("\n- Hashing training script...")
		scriptHash, err := hashFile(scriptFile)
		if err != nil {
			fmt.Printf("\n Error hashing script file: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("  - Script Hash: %s\n", scriptHash)

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

		// Create hidden directory
		hiddenDir := ".trustmint"
		if err := os.MkdirAll(hiddenDir, 0755); err != nil {
			fmt.Printf("\n Error creating hidden directory: %s\n", err)
			os.Exit(1)
		}

		viper.Set("config_hash", configHash)
		viper.Set("script_hash", scriptHash) // Add script hash
		viper.Set("dataset_hash", datasetHash)
		viper.Set("model_hash", modelHash)
		hashesFile := filepath.Join(hiddenDir, ".hashes.yaml")
		if err := viper.WriteConfigAs(hashesFile); err != nil {
			fmt.Printf("\n Error saving hashes: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("\n- Proofs saved to %s\n", hashesFile)

		// 3. Sign the findings
		fmt.Println("\n- Signing artifacts...")
		if AppPrivateKey == "" {
			fmt.Println("⚠️  WARNING: No Private Key found. Skipping signature.")
			fmt.Println("   (This is expected in dev mode if -ldflags was not used)")
		} else {
			// Include scriptHash in signature
			signature, err := signData(configHash + datasetHash + modelHash + scriptHash)
			if err != nil {
				fmt.Printf("\n Error generating signature: %s\n", err)
				os.Exit(1)
			}
			sigFile := filepath.Join(hiddenDir, ".hashes.sig")
			if err := os.WriteFile(sigFile, []byte(signature), 0644); err != nil {
				fmt.Printf("\n Error saving signature: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("  - Signature generated and saved to %s\n", sigFile)
		}

		fmt.Println("\n\n--- ✅ Secure Training Process Finished ---")
	},
}

// hashDirectory function remains the same...
// hashDirectory function remains the same...
func hashDirectory(path string) (string, error) {
	hasher := sha256.New()
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := io.Copy(hasher, file); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func hashFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func signData(data string) (string, error) {
	block, _ := pem.Decode([]byte(AppPrivateKey))
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block containing the key")
	}
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256([]byte(data))
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}

	// Use ASN.1 DER encoding for standard compatibility
	type ECDSASignature struct {
		R, S *big.Int
	}
	signature, err := asn1.Marshal(ECDSASignature{R: r, S: s})
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(signature), nil
}

func init() {
	rootCmd.AddCommand(trainCmd)
}
