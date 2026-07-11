package notification

import (
	"os"

	"github.com/fehmicorp/agent/v1/res/assets"
)

func tempIcon(iconPath string) (string, error) {
	if iconPath == "" {
		return "", nil
	}

	data, err := assets.Read(iconPath)
	if err != nil {
		return "", err
	}

	tmp, err := os.CreateTemp("", "fehmi-*")
	if err != nil {
		return "", err
	}

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		os.Remove(tmp.Name())
		return "", err
	}

	tmp.Close()

	return tmp.Name(), nil
}
