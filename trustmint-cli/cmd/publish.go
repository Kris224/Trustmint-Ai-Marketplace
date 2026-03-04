package cmd

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// PublishResponse is the JSON returned by the backend /publish endpoint
type PublishResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	TokenID   int    `json:"token_id"`
	ListingID int    `json:"listing_id"`
	TxHash    string `json:"tx_hash"`
	IpfsCid   string `json:"ipfs_cid"`
}

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publishes the model, mints an NFT, and lists it on the marketplace.",
	Run: func(cmd *cobra.Command, args []string) {
		printHeader("Stage 2: ONLINE Publishing")

		// Check wallet binding
		if WalletAddress == "" {
			printError("No wallet address found in this CLI binary.")
			fmt.Println("   Please download a fresh CLI from the Developer Dashboard.")
			return
		}
		fmt.Printf("🔑 Wallet: %s\n\n", WalletAddress)

		fmt.Println("▶️  1/4: Loading cryptographic proofs...")
		viper.SetConfigFile(".trustmint/.hashes.yaml")
		if err := viper.ReadInConfig(); err != nil {
			printError("Could not find proofs file. Run 'trustmint train' first.")
			return
		}
		fmt.Printf("✅ Proofs loaded.\n\n")

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		// Send hashes and wallet address
		writer.WriteField("wallet_address", WalletAddress)
		writer.WriteField("config_hash", viper.GetString("config_hash"))
		writer.WriteField("dataset_hash", viper.GetString("dataset_hash"))
		writer.WriteField("model_hash", viper.GetString("model_hash"))
		writer.WriteField("script_hash", viper.GetString("script_hash"))

		// Read and send signature
		sigBytes, err := os.ReadFile(".trustmint/.hashes.sig")
		if err == nil {
			writer.WriteField("signature_hex", string(sigBytes))
		} else {
			fmt.Println("⚠️  Warning: .hashes.sig not found. Uploading without signature.")
		}

		fmt.Println("▶️  2/4: Attaching files for upload...")
		addFileToRequest(writer, "config_file", "trustmint.yml")
		viper.SetConfigFile("trustmint.yml")
		viper.ReadInConfig()
		modelPath := filepath.Join(viper.GetString("model_output_dir"), "model.pkl")
		addFileToRequest(writer, "model_file", modelPath)
		addFileToRequest(writer, "script_file", "train.py")
		fmt.Printf("✅ Config, model, and script files attached.\n\n")

		fmt.Println("▶️  3/4: Compressing and attaching dataset...")
		datasetDir := viper.GetString("dataset_dir")
		datasetZipPath := "dataset.zip"

		if err := zipSource(datasetDir, datasetZipPath); err != nil {
			printError(fmt.Sprintf("Failed to zip dataset: %v", err))
			return
		}
		addFileToRequest(writer, "dataset_zip", datasetZipPath)
		defer os.Remove(datasetZipPath)
		fmt.Printf("✅ Dataset attached as dataset.zip.\n\n")

		writer.Close()

		fmt.Println("▶️  4/4: Publishing to backend (IPFS + blockchain)...")
		url := BackendURL + "/publish"
		req, _ := http.NewRequest("POST", url, body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		fmt.Printf("   - Destination: %s\n", url)
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			printError(fmt.Sprintf("Connection to backend failed: %s", err))
			return
		}
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)

		if resp.StatusCode == 200 {
			var result PublishResponse
			if err := json.Unmarshal(respBody, &result); err == nil && result.TokenID > 0 {
				fmt.Println()
				fmt.Println("🎉 ══════════════════════════════════════════════")
				fmt.Println("   MODEL PUBLISHED SUCCESSFULLY!")
				fmt.Println("   ══════════════════════════════════════════════")
				fmt.Printf("   🪙 NFT Token ID  : #%d\n", result.TokenID)
				fmt.Printf("   🛒 Listing ID    : #%d\n", result.ListingID)
				fmt.Printf("   🔗 Tx Hash       : %s\n", result.TxHash)
				fmt.Printf("   📦 IPFS CID      : %s\n", result.IpfsCid)
				fmt.Printf("   👛 Owner Wallet  : %s\n", WalletAddress)
				fmt.Println("   ══════════════════════════════════════════════")
				fmt.Println("   Your model is now visible on the marketplace!")
			} else {
				printSuccess("All artifacts have been published and saved by the backend.")
			}
		} else {
			printError(fmt.Sprintf("Backend rejected the submission. Status: %s\nBody: %s", resp.Status, string(respBody)))
		}
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}

// --- Helper Functions ---
func addFileToRequest(writer *multipart.Writer, fieldname, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		printError(fmt.Sprintf("Could not open file %s: %v", filename, err))
		return
	}
	defer file.Close()
	part, _ := writer.CreateFormFile(fieldname, filepath.Base(filename))
	io.Copy(part, file)
}

func zipSource(source, target string) error {
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()
	writer := zip.NewWriter(f)
	defer writer.Close()

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Method = zip.Deflate
		header.Name, err = filepath.Rel(source, path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(headerWriter, f)
		return err
	})
}
