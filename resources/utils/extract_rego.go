package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func extractRegoFiles(tempPACPath string, taxons []string) (string, error) {
	// Create a directory to store all .rego files
	outputDir := filepath.Join(tempPACPath, "rego_files")
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create output directory: %v", err)
	}

	// Iterate over each taxon
	for _, taxon := range taxons {
		// Use filepath.Join to handle spaces and other path issues
		taxonPath := filepath.Join(tempPACPath, taxon)

		// Check if the taxon directory exists
		if _, err := os.Stat(taxonPath); os.IsNotExist(err) {
			fmt.Printf("Directory %s does not exist, skipping...\n", taxonPath)
			continue
		}

		// Walk through the directory
		err := filepath.Walk(taxonPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Check if the file has a .rego extension
			if !info.IsDir() && filepath.Ext(path) == ".rego" {
				// Copy the .rego file to the output directory
				destPath := filepath.Join(outputDir, info.Name())
				err := copyFile(path, destPath)
				if err != nil {
					return fmt.Errorf("failed to copy file %s: %v", path, err)
				}
			}
			return nil
		})

		if err != nil {
			return "", fmt.Errorf("error walking the path %s: %v", taxonPath, err)
		}
	}

	return outputDir, nil
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
