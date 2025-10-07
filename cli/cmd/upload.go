package cmd

import (
	"fmt"

	"github.com/Niyalurahman/trustmint/internal/upload"
	"github.com/spf13/cobra"
)

var metadataFile string

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload model artifacts and mint ownership (placeholder)",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Basic validation
		if metadataFile == "" {
			return fmt.Errorf("metadata file is required --metadata <path>")
		}
		// Call placeholder upload
		if err := upload.UploadModel(metadataFile); err != nil {
			return err
		}
		fmt.Println("✅ Upload placeholder completed for:", metadataFile)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringVar(&metadataFile, "metadata", "output/metadata.json", "Path to metadata JSON produced by train")
}
