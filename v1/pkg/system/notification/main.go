package notification

import (
	"fmt"
	"runtime"
)

var provider Provider

func Register(appID string) error {

	switch runtime.GOOS {

	case "windows":
		provider = &windowsProvider{}

	case "darwin":
		provider = &darwinProvider{}

	case "linux":
		provider = &linuxProvider{}

	default:
		return fmt.Errorf("notifications not supported on %s", runtime.GOOS)
	}

	return provider.Register(appID)
}

func Push(opt *NotificationOptions) error {

	if provider == nil {
		if err := Register(opt.AppID); err != nil {
			return err
		}
	}

	return provider.Push(opt)
}

func Close() error {

	if provider == nil {
		return nil
	}

	return provider.Close()
}
