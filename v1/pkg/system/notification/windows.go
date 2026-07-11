package notification

import (
	"os"

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

	if path, err := tempIcon(opt.IconPath); err == nil {
		defer os.Remove(path)
		n.Icon = path
	}

	return n.Push()
}
