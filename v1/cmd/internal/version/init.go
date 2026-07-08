package version

type Info struct {
	Version   string
	BuildType string
	BuildDate string
	Commit    string
	Branch    string
}

var Current = &Info{
	Version:   "0.0.1",
	BuildType: "development",
	BuildDate: "",
	Commit:    "local",
	Branch:    "main",
}
