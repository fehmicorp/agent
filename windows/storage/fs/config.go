package fs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Standard Base Path Topography Structure Configurations
var (
	BaseDir    = "C:\\ProgramData\\Fehmi"
	BaseData   = BaseDir + "\\agent\\data"
	BaseConfig = BaseDir + "\\agent\\config"
	BaseLogs   = BaseDir + "\\agent\\logs"
)

// EnsureDir checks if a directory path exists; if it doesn't, it creates it with secure permissions (0755).
func EnsureDir(dirPath string) (bool, error) {
	if dirPath == "" {
		return false, fmt.Errorf("directory path cannot be empty")
	}

	info, err := os.Stat(dirPath)
	if err == nil {
		if info.IsDir() {
			return false, nil // Exists and is already a directory
		}
		return false, fmt.Errorf("path exists but is a file, not a directory: %s", dirPath)
	}

	if os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			return false, fmt.Errorf("failed to create directory structure: %w", err)
		}
		return true, nil // Successfully created
	}

	return false, err
}

// InitSubsystems provisions the entire application path ecosystem in one go
func InitSubsystems() error {
	paths := []string{BaseDir, BaseData, BaseConfig, BaseLogs}
	for _, p := range paths {
		if _, err := EnsureDir(p); err != nil {
			return fmt.Errorf("failed bootstrapping workspace directory %s: %w", p, err)
		}
	}
	return nil
}

// WriteFile writes string or byte data to a file safely. It automatically creates the parent folder if missing.
func WriteFile(filePath string, data []byte) error {
	dir := filepath.Dir(filePath)
	if _, err := EnsureDir(dir); err != nil {
		return err
	}

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

// DeletePath removes a target file or an entire directory recursively.
func DeletePath(path string) error {
	if !Exists(path) {
		return nil // Already absent
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

// CopyFile duplicates a file from a source location to a destination location safely.
func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

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
		return fmt.Errorf("failed streaming file transfer: %w", err)
	}
	return destFile.Sync()
}
