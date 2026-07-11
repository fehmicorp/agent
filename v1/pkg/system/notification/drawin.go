package notification

import (
	"os"

	"github.com/gen2brain/beeep"
)

type darwinProvider struct {
	appID string
}

func (d *darwinProvider) Register(appID string) error {
	d.appID = appID
	return nil
}

func (d *darwinProvider) Close() error {
	return nil
}

func (d *darwinProvider) Push(opt *NotificationOptions) error {

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
