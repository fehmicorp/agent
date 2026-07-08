package appinfo

func DevelopmentConfig() {

	Current.BuildType = Development

	Current.Debug = true
	Current.Console = true
	Current.DevTools = true
	Current.Admin = false
	Current.Profiling = true
}
