package notification

import "github.com/gen2brain/beeep"

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

	return beeep.Notify(
		opt.Title,
		opt.Message,
		opt.IconPath,
	)
}
