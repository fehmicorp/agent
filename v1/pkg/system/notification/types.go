package notification

type NotificationOptions struct {
	AppID    string
	Title    string
	Message  string
	Subtitle string
	Body     string
	IconPath string
	Sound    bool
	Silent   bool

	Actions []Action
}

type Action struct {
	ID    string
	Title string
}

type Provider interface {
	Register(appID string) error
	Push(*NotificationOptions) error
	Close() error
}
