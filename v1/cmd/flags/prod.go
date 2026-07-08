package flags

func ProductionConfig() {

	Current.BuildType = Production

	Current.Debug = false
	Current.Console = false
	Current.DevTools = false
	Current.Admin = true
	Current.Profiling = false
}
