package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"io/fs"
	"regexp"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func getTaxons(resources []string) map[string]string {
	taxons := map[string]string{}
	for _, resource := range resources {
		tfName, ok := TerraformResourceToTerraformName[resource]
		if ok {
			if taxon, ok := terraformToTaxonName[tfName]; ok {
				taxons[taxon.CloudServiceName] = "DEFAULT"
			}
		}
	}
	return taxons
}

func getResources(iacPath string) ([]string, error) {
	// Set to store unique resource types
	resourceTypes := make(map[string]struct{})

	// Regex to match resource blocks in Terraform files
	resourceRegex := regexp.MustCompile(`resource\s+"([^"]+)"\s+"([^"]+)"`)

	// Walk through all files in the directory
	err := filepath.Walk(iacPath, func(path string, info fs.FileInfo, err error) error {
		// Check for errors during walking
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Only process .tf files
		if !strings.HasSuffix(path, ".tf") {
			return nil
		}

		// Read file contents
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading file %s: %v", path, err)
		}

		// Find all resource matches in the file
		matches := resourceRegex.FindAllStringSubmatch(string(content), -1)
		for _, match := range matches {
			if len(match) >= 2 {
				resourceTypes[match[1]] = struct{}{}
			}
		}

		return nil
	})

	// Check for errors during file walking
	if err != nil {
		return nil, fmt.Errorf("error walking directory: %v", err)
	}

	// Convert map keys to slice
	result := make([]string, 0, len(resourceTypes))
	for resourceType := range resourceTypes {
		result = append(result, resourceType)
	}

	return result, nil
}

func getTaxonsByIAC(iacPath string) ([]string, error) {
	resources, err := getResources(iacPath)
	if err != nil {
		return nil, err
	}

	result := []string{}
	taxons := getTaxons(resources)
	for taxon := range taxons {
		result = append(result, taxon)
	}
	return result, nil
}

func GetDefaultPAC(iacPath, pacVersion string) (string, error) {
	taxons, err := getTaxonsByIAC(iacPath)
	if err != nil {
		return "", err
	}

	tempCloneDir, err := os.MkdirTemp("", "pac-clone-*")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary directory: %v", err)
	}

	tempPACPath, err := getPACPath(tempCloneDir, pacVersion, taxons)
	if err != nil {
		return "", err
	}

	relevantPACPath, err := extractRegoFiles(tempPACPath, taxons)
	if err != nil {
		return "", err
	}
	return relevantPACPath, nil
}

func getPACPath(tempDir string, branch string, taxons []string) (string, error) {
	repoURL := "https://github.com/nonfx/starchitect-cloudguard"
	folderPath := "terraform/aws"
	if branch == "" {
		branch = "main"
	}

	_, err := git.PlainClone(tempDir, false, &git.CloneOptions{
		URL:           repoURL,
		Progress:      os.Stdout,
		Depth:         1,
		ReferenceName: plumbing.NewBranchReferenceName(branch),
	})
	if err != nil {
		return "", fmt.Errorf("failed to clone rules repository: %v", err)
	}

	// Step 3: Construct the path to the desired folder
	clonedFolderPath := filepath.Join(tempDir, folderPath)
	if _, err := os.Stat(clonedFolderPath); os.IsNotExist(err) {
		return "", fmt.Errorf("folder %s does not exist in the cloned repository", clonedFolderPath)
	}
	return clonedFolderPath, nil
}
