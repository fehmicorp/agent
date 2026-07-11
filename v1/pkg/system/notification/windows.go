package notification

import (
	"path/filepath"

	"github.com/go-toast/toast"
)

type windowsProvider struct {
	appID string
}

func (w *windowsProvider) Register(appID string) error {
	w.appID = appID
	return nil
}

func (w *windowsProvider) Close() error {
	return nil
}

func (w *windowsProvider) Push(opt *NotificationOptions) error {

	n := toast.Notification{
		AppID:   w.appID,
		Title:   opt.Title,
		Message: opt.Message,
	}

	if opt.IconPath != "" {
		if abs, err := filepath.Abs(opt.IconPath); err == nil {
			n.Icon = abs
		}
	}

	return n.Push()
}
