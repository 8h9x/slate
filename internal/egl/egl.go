package egl

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func LocateInstallation(appName string) (string, error) {
	bytes, err := os.ReadFile(path.Join(os.Getenv("PROGRAMDATA"), "Epic/UnrealEngineLauncher/LauncherInstalled.dat"))
	if err != nil {
		return "", err
	}

	var data EpicLauncherInstallData
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return "", err
	}

	for _, entry := range data.InstallationList {
		if entry.AppName == appName {
			return entry.InstallLocation, nil
		}
	}

	return "", fmt.Errorf("no '%s' app found in epic games launcher install manifest", appName)
}
