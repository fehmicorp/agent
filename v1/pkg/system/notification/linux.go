package notification

import "github.com/gen2brain/beeep"

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

	return beeep.Notify(
		opt.Title,
		opt.Message,
		opt.IconPath,
	)
}
