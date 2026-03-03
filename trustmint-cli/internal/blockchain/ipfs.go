package blockchain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// PinataConfig holds Pinata API configuration
type PinataConfig struct {
	APIKey    string
	SecretKey string
	JWT       string
}

// IPFSResponse represents the response from Pinata after pinning
type IPFSResponse struct {
	IpfsHash  string `json:"IpfsHash"`
	PinSize   int    `json:"PinSize"`
	Timestamp string `json:"Timestamp"`
}

// UploadToPinata uploads a file to IPFS via Pinata
func UploadToPinata(filePath string, config PinataConfig) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add file
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %w", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return "", fmt.Errorf("failed to copy file: %w", err)
	}

	// Add metadata

	pinataMetadata := map[string]interface{}{
		"name": filepath.Base(filePath),
	}
	pinataMetadataBytes, _ := json.Marshal(pinataMetadata)

	_ = writer.WriteField("pinataMetadata", string(pinataMetadataBytes))
	_ = writer.WriteField("pinataOptions", `{"cidVersion": 1}`)

	writer.Close()

	// Create request
	url := "https://api.pinata.cloud/pinning/pinFileToIPFS"
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if config.JWT != "" {
		req.Header.Set("Authorization", "Bearer "+config.JWT)
	} else {
		req.Header.Set("pinata_api_key", config.APIKey)
		req.Header.Set("pinata_secret_api_key", config.SecretKey)
	}

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Parse response
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("pinata API error (%d): %s", resp.StatusCode, string(bodyBytes))
	}

	var ipfsResp IPFSResponse
	if err := json.NewDecoder(resp.Body).Decode(&ipfsResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return ipfsResp.IpfsHash, nil
}

// UploadJSONToPinata uploads JSON data to IPFS via Pinata
func UploadJSONToPinata(data interface{}, name string, config PinataConfig) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Create request body
	reqBody := map[string]interface{}{
		"pinataContent": json.RawMessage(jsonData),
		"pinataMetadata": map[string]string{
			"name": name,
		},
		"pinataOptions": map[string]interface{}{
			"cidVersion": 1,
		},
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create request
	url := "https://api.pinata.cloud/pinning/pinJSONToIPFS"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	if config.JWT != "" {
		req.Header.Set("Authorization", "Bearer "+config.JWT)
	} else {
		req.Header.Set("pinata_api_key", config.APIKey)
		req.Header.Set("pinata_secret_api_key", config.SecretKey)
	}

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Parse response
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("pinata API error (%d): %s", resp.StatusCode, string(bodyBytes))
	}

	var ipfsResp IPFSResponse
	if err := json.NewDecoder(resp.Body).Decode(&ipfsResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return ipfsResp.IpfsHash, nil
}
