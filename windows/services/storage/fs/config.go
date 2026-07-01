package fs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// EnsureDir checks if a directory path exists; if it doesn't, it creates it with secure permissions (0755).
func EnsureDir(dirPath string) (bool, error) {
	if dirPath == "" {
		return false, fmt.Errorf("directory path cannot be empty")
	}

	// Check if directory already exists
	info, err := os.Stat(dirPath)
	if err == nil {
		if info.IsDir() {
			return false, nil // Exists and is already a directory
		}
		return false, fmt.Errorf("path exists but is a file, not a directory: %s", dirPath)
	}

	// Create directory and any missing parent directories
	if os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			return false, fmt.Errorf("failed to create directory structure: %w", err)
		}
		return true, nil // Successfully created
	}

	return false, err
}

// WriteFile writes string or byte data to a file safely. It automatically creates the parent folder if missing.
func WriteFile(filePath string, data []byte) error {
	dir := filepath.Dir(filePath)
	if _, err := EnsureDir(dir); err != nil {
		return err
	}

	// Write file atomically (0644 gives read/write permissions to owner, read to others)
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file content: %w", err)
	}
	return nil
}

// ReadFile reads and returns the complete contents of a local file.
func ReadFile(filePath string) ([]byte, error) {
	if !Exists(filePath) {
		return nil, os.ErrNotExist
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %w", err)
	}
	return data, nil
}

// DeletePath removes a target file or an entire directory (along with all its contents).
func DeletePath(path string) error {
	if !Exists(path) {
		return nil // Already gone
	}

	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("failed to delete path %s: %w", path, err)
	}
	return nil
}

// Exists is a quick helper to verify if a file or folder exists on disk.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// GetUserConfigDir returns the standard application profile directory path for FehmiAgent (%APPDATA%\FehmiAgent on Windows).
func GetAgentDataDir() string {
	baseDir, err := os.UserConfigDir()
	if err != nil {
		// Fallback to local execution directory if OS environment is constrained
		baseDir = "."
	}
	return filepath.Join(baseDir, "FehmiAgent")
}

// CopyFile duplicates a file from a source location to a destination location safely.
func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Ensure destination path folder is provisioned
	destDir := filepath.Dir(dst)
	if _, err := EnsureDir(destDir); err != nil {
		return err
	}

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed during streaming file copy block copy: %w", err)
	}
	return destFile.Sync()
}
