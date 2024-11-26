package resources

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func getPACPath(tempDir string) (string, error) {
	repoURL := "https://github.com/nonfx/starchitect-cloudguard"
	folderPath := "terraform/aws"
	branch := "main"

	_, err := git.PlainClone(tempDir, false, &git.CloneOptions{
		URL:           repoURL,
		Progress:      os.Stdout,
		Depth:         1, // Clone only the latest commit. TODO: accept version.
		ReferenceName: plumbing.NewBranchReferenceName(branch),
	})
	if err != nil {
		return "", fmt.Errorf("failed to clone repository: %v", err)
	}

	// Step 3: Construct the path to the desired folder
	clonedFolderPath := filepath.Join(tempDir, folderPath)
	if _, err := os.Stat(clonedFolderPath); os.IsNotExist(err) {
		return "", fmt.Errorf("folder %s does not exist in the cloned repository", clonedFolderPath)
	}
	return clonedFolderPath, nil
}
