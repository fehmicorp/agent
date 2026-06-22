package config

func Get() *App {
	return instance
}

func GetTray() *Tray {

	if instance == nil {
		return nil
	}

	return &instance.Tray
}

func GetFunctions() []Functions {

	if instance == nil {
		return nil
	}

	return instance.Functions
}
