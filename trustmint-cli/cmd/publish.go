package cmd

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publishes the model, dataset, and proofs to the local backend.",
	Run: func(cmd *cobra.Command, args []string) {
		printHeader("Stage 2: ONLINE Publishing")

		fmt.Println("▶️  1/4: Loading cryptographic proofs...")
		viper.SetConfigFile(".hashes.yaml")
		if err := viper.ReadInConfig(); err != nil {
			printError("Could not find proofs file. Run 'trustmint train' first.")
			return
		}
		fmt.Printf("✅ Proofs loaded.\n\n")

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		writer.WriteField("config_hash", viper.GetString("config_hash"))
		writer.WriteField("dataset_hash", viper.GetString("dataset_hash"))
		writer.WriteField("model_hash", viper.GetString("model_hash"))

		fmt.Println("▶️  2/4: Attaching files for upload...")
		addFileToRequest(writer, "config_file", "trustmint.yml")
		viper.SetConfigFile("trustmint.yml")
		viper.ReadInConfig()
		modelPath := filepath.Join(viper.GetString("model_output_dir"), "model.pkl")
		addFileToRequest(writer, "model_file", modelPath)
		fmt.Printf("✅ Config and model files attached.\n\n")
		
		fmt.Println("▶️  3/4: Compressing and attaching dataset...")
		datasetDir := viper.GetString("dataset_dir")
		datasetZipPath := "dataset.zip"
		if err := zipSource(datasetDir, datasetZipPath); err != nil {
			printError(fmt.Sprintf("Failed to zip dataset: %v", err))
			return
		}
		addFileToRequest(writer, "dataset_zip", datasetZipPath)
		defer os.Remove(datasetZipPath) // Clean up the zip file after sending
		fmt.Printf("✅ Dataset attached as dataset.zip.\n\n")
		
		writer.Close()

		fmt.Println("▶️  4/4: Publishing all artifacts to local backend...")
		url := "http://127.0.0.1:5001/publish"
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

		if resp.StatusCode == 200 {
			printSuccess("All artifacts have been published and saved by the backend.")
		} else {
			printError(fmt.Sprintf("Backend rejected the submission. Status: %s", resp.Status))
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
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
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