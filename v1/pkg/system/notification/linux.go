package notification

import (
	"os"

	"github.com/gen2brain/beeep"
)

type linuxProvider struct {
	appID string
}

func (l *linuxProvider) Register(appID string) error {
	l.appID = appID
	return nil
}

func (l *linuxProvider) Close() error {
	return nil
}

func (l *linuxProvider) Push(opt *NotificationOptions) error {

	icon := ""

	if path, err := tempIcon(opt.IconPath); err == nil {
		defer os.Remove(path)
		icon = path
	}

	return beeep.Notify(
		opt.Title,
		opt.Message,
		icon,
	)
}
