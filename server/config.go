package server

import (
	"fmt"
	"os"
)

type WFConfig struct {
	Port                   string
	IndexFilepath          string
	ShareDirectoryFilepath string
}

var DefaultUnixConfig = WFConfig{
	Port:                   "1997",
	IndexFilepath:          "wi-fileshare/index.md",
	ShareDirectoryFilepath: "wi-fileshare/share",
}

var DefaultWindowsConfig = WFConfig{
	Port:                   "1997",
	IndexFilepath:          "wi-fileshare\\index.md",
	ShareDirectoryFilepath: "wi-fileshare\\share",
}

func validateConfig(config WFConfig) (WFConfig, error) {
	validatedShareDirectoryFilepath, err := validateShareDirectoryFilepath(config.ShareDirectoryFilepath)
	if err != nil {
		return WFConfig{}, err
	}

	validatedIndexFilepath, err := validateIndexFilepath(config.IndexFilepath)
	if err != nil {
		return WFConfig{}, err
	}

	validatedConfig := WFConfig {
		Port: config.Port,
		IndexFilepath: validatedIndexFilepath,
		ShareDirectoryFilepath: validatedShareDirectoryFilepath,
	}

	return validatedConfig, nil
}

func validateIndexFilepath(filePath string) (string, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		f, err := os.Create(filePath)
		f.WriteString("# index\n[**shared files**](/share/)")
		if err != nil {
			return "filePath", err
		}
		defer f.Close()
	}

	return filePath, nil
}

func validateShareDirectoryFilepath(filePath string) (string, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) || os.IsPermission(err) {
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return filePath, err
		}
	}

	if fileInfo, _ := os.Stat(filePath); fileInfo.IsDir() == false {
		return filePath, fmt.Errorf("invalid share directory: %v is not a directory", filePath)
	}

	return filePath, nil
}
