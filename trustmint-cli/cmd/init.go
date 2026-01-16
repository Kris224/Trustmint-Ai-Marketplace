package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new Trustmint project.",
	Long:  `Generates the necessary configuration files and directory structure for a new Trustmint project. If a project name is provided, it creates a new directory with that name.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("✨ Initializing new Trustmint project...")

		// Determine project directory
		projectDir := "trustmint-project"
		if len(args) > 0 {
			projectDir = args[0]
		}

		// Create main project directory
		if err := os.MkdirAll(projectDir, 0755); err != nil {
			fmt.Printf("❌ Human error: Failed to create project directory %s: %v\n", projectDir, err)
			os.Exit(1)
		}
		fmt.Printf("   - Created project directory: %s/\n", projectDir)

		// 1. Create Subdirectories (dataset, output)
		subDirs := []string{"dataset", "output"}
		for _, dir := range subDirs {
			fullPath := filepath.Join(projectDir, dir)
			if err := os.MkdirAll(fullPath, 0755); err != nil {
				fmt.Printf("❌ Human error: Failed to create directory %s: %v\n", fullPath, err)
				os.Exit(1)
			}
			fmt.Printf("   - Created directory: %s/%s/\n", projectDir, dir)
		}

		// 2. Create trustmint.yml
		configContent := `dataset_dir: ./dataset
model_output_dir: ./output
training_script: python3 train.py
environment:
  network: none
  read_only_root: false
  volumes:
    - source: ./dataset
      target: /app/dataset
      read_only: true
    - source: .
      target: /app/project
      read_only: false
`
		configFile := filepath.Join(projectDir, "trustmint.yml")
		if err := os.WriteFile(configFile, []byte(configContent), 0644); err != nil {
			fmt.Printf("❌ Human error: Failed to create trustmint.yml: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("   - Created config: %s\n", configFile)

		// 3. Create train.py template
		trainContent := `import os
import time

print("--- Python training script started ---")

dataset_path = './dataset/data.csv'
model_output_path = './output/model.pkl'

try:
    with open(dataset_path, 'r') as f:
        lines = f.readlines()
        print(f"   - Successfully read {len(lines)} lines from the dataset.")
except Exception as e:
    print(f"   - Error reading dataset: {e}")
    exit(1)

print("   - Simulating model training for 5 seconds...")
# time.sleep(5) 

try:
    with open(model_output_path, 'w') as f:
        f.write("This is a dummy trained model.")
    print(f"   - Model saved successfully to {model_output_path}")
except Exception as e:
    print(f"   - Error saving model: {e}")
    exit(1)

print("--- Python training script finished ---")
`
		trainFile := filepath.Join(projectDir, "train.py")
		if err := os.WriteFile(trainFile, []byte(trainContent), 0644); err != nil {
			fmt.Printf("❌ Human error: Failed to create train.py: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("   - Created template: %s\n", trainFile)

		// 4. Create dummy csv in dataset to avoid errors
		datasetFile := filepath.Join(projectDir, "dataset", "data.csv")
		if err := os.WriteFile(datasetFile, []byte("id,value\n1,100"), 0644); err != nil {
			fmt.Printf("❌ Human error: Failed to create dummy dataset: %v\n", err)
		}
		fmt.Printf("   - Created dummy dataset: %s\n", datasetFile)

		fmt.Println("\n✅ Project initialized successfully!")
		fmt.Printf("   cd %s\n", projectDir)
		fmt.Println("   trustmint start")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
