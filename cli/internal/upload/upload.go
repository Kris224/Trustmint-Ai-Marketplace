package upload

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Niyalurahman/trustmint/internal/config"
	"github.com/Niyalurahman/trustmint/internal/chain"
)

// ipfsAddFile uploads a single file to a local IPFS daemon HTTP API and returns the CID.
func ipfsAddFile(apiURL string, filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(part, f); err != nil {
		return "", err
	}
	if err := writer.Close(); err != nil {
		return "", err
	}

	// POST to /api/v0/add
	req, err := http.NewRequest("POST", apiURL+"/api/v0/add?pin=true", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		slurp, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ipfs add failed: %s", string(slurp))
	}

	var out struct {
		Name string `json:"Name"`
		Hash string `json:"Hash"`
		Size string `json:"Size"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	return out.Hash, nil
}

// UploadModel uploads metadata and artifacts to configured off-chain storage (IPFS when configured).
// - metadataPath: path to metadata.json
// It will read trustmint.yaml from current working directory for storage preferences if available.
func UploadModel(metadataPath string) error {
	if _, err := os.Stat(metadataPath); os.IsNotExist(err) {
		return fmt.Errorf("metadata file does not exist: %s", metadataPath)
	}

	// Try to load trustmint.yaml if present
	var cfg *config.Config
	if _, err := os.Stat("trustmint.yaml"); err == nil {
		c, err := config.LoadConfig("trustmint.yaml")
		if err == nil {
			cfg = c
		}
	}

	storage := "local"
	if cfg != nil && cfg.Storage != "" {
		storage = cfg.Storage
	}

	if storage == "ipfs" {
		apiURL := "http://127.0.0.1:5001"
		fmt.Println("🔗 Detected IPFS storage preference. Attempting to upload files to IPFS at", apiURL)
		cid, err := ipfsAddFile(apiURL, metadataPath)
		if err != nil {
			return fmt.Errorf("failed to upload metadata to IPFS: %w", err)
		}
		fmt.Println("✅ Uploaded metadata to IPFS with CID:", cid)

		// Attempt on-chain mint (placeholder)
		tokenID, mErr := chain.MintNFT(cid)
		if mErr != nil {
			fmt.Println("Warning: on-chain mint (placeholder) failed:", mErr)
		} else {
			fmt.Println("✅ (placeholder) Minted NFT token ID:", tokenID)
		}

		// Upload other files in same directory as metadata
		dir := filepath.Dir(metadataPath)
		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if info.IsDir() {
				return nil
			}
			// skip metadata file itself (already uploaded)
			if filepath.Base(path) == filepath.Base(metadataPath) {
				return nil
			}
			fmt.Println("Uploading", path)
			if c, err := ipfsAddFile(apiURL, path); err == nil {
				fmt.Println(" -> CID:", c)
			} else {
				fmt.Println(" -> failed:", err)
			}
			return nil
		})
		if err != nil {
			fmt.Println("Warning: error while uploading additional files:", err)
		}
		return nil
	}

	// Default: local storage (no-op)
	fmt.Println("Storage is set to", storage, "- no upload performed (local only).")
	return nil
}
