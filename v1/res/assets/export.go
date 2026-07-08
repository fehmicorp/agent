package assets

import (
	"os"
	"path/filepath"
)

func Export(name, destination string) error {

	data, err := Read(name)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(destination), 0755); err != nil {
		return err
	}

	return os.WriteFile(destination, data, 0644)
}
