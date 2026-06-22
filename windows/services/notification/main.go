package notification

import (
	"fmt"
	"path/filepath"

	"github.com/fehmicorp/agent/windows/config/types"
	"github.com/go-toast/toast"
)

func Push(opt *types.NotiOptions) error {
	notification := toast.Notification{
		AppID:   "Fehmi Endpoint Agent",
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
