package appinfo

type AppInfo struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Build       string `json:"build"`
	BuildType   string `json:"buildType"`
	Company     string `json:"company"`
	Website     string `json:"website"`
	Endpoint    string `json:"endpoint"`
	Description string `json:"description"`
}
