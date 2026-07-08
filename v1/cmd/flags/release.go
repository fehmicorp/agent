package internal

func ReleaseConfig() {

	Current.BuildType = Release

	Current.Debug = false
	Current.Console = false
	Current.DevTools = false
	Current.Admin = true
	Current.Profiling = false
}
