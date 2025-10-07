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
)

type PinataResp struct {
	IpfsHash  string `json:"IpfsHash"`
	PinSize   int    `json:"PinSize"`
	Timestamp string `json:"Timestamp"`
}

// UploadFileToPinata uploads a single file and returns IpfsHash
func UploadFileToPinata(jwt, filePath string) (string, error) {
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
	writer.Close()

	req, err := http.NewRequest("POST", "https://api.pinata.cloud/pinning/pinFileToIPFS", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+jwt)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		b, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("pinata upload failed: %s", string(b))
	}

	var pr PinataResp
	if err := json.NewDecoder(resp.Body).Decode(&pr); err != nil {
		return "", err
	}
	return pr.IpfsHash, nil
}
