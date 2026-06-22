package notification

import (
	"fmt"
	"path/filepath"

	"github.com/go-toast/toast"
)

type Options struct {
	Title    string
	Message  string
	IconPath string
}

func Push(opt Options) error {
	notification := toast.Notification{
		AppID:   "Fehmi Agent",
		Title:   opt.Title,
		Message: opt.Message,
	}
	if opt.IconPath != "" {
		absPath, err := filepath.Abs(opt.IconPath)
		if err == nil {
			notification.Icon = absPath
		}
	}
	err := notification.Push()
	if err != nil {
		return fmt.Errorf("failed to dispatch toast notification: %w", err)
	}
	return nil
}
