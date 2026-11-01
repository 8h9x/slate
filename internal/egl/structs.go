package egl

type EpicLauncherInstallData struct {
	InstallationList []EpicLauncherInstallDataInstallation `json:"InstallationList"`
}

type EpicLauncherInstallDataInstallation struct {
	InstallLocation string `json:"InstallLocation"`
	NamespaceID     string `json:"NamespaceId"`
	ItemID          string `json:"ItemId"`
	ArtifactID      string `json:"ArtifactId"`
	AppVersion      string `json:"AppVersion"`
	AppName         string `json:"AppName"`
}
