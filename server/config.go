package server

import (
	"fmt"
	"os"
	"strconv"
)

type WFConfig struct {
	Port                   string
	IndexFilepath          string
	ShareDirectoryFilepath string
}

var DefaultConfig = WFConfig{
	Port:                   "8080",
	IndexFilepath:          "index.md",
	ShareDirectoryFilepath: "share/",
}

func validateConfig(config WFConfig) (WFConfig, error) {
	validatedPort, err := validatePort(config.Port)
	if err != nil {
		return DefaultConfig, err
	}

	validatedIndexFilepath, err := validateIndexFilepath(config.IndexFilepath)
	if err != nil {
		return DefaultConfig, err
	}

	validatedShareDirectoryFilepath, err := validateShareDirectoryFilepath(config.ShareDirectoryFilepath)
	if err != nil {
		return DefaultConfig, err
	}

	validatedConfig := WFConfig {
		Port: validatedPort,
		IndexFilepath: validatedIndexFilepath,
		ShareDirectoryFilepath: validatedShareDirectoryFilepath,
	}

	return validatedConfig, nil
}

func validatePort(input string) (string, error) {
	num, err := strconv.Atoi(input)
	if err != nil {
		return DefaultConfig.Port, fmt.Errorf("invalid port: %v", err)
	}

	if num < 0 || num > 65536 {
		return DefaultConfig.Port, fmt.Errorf("invalid port: %d is not between 0 and 65536", num)
	}

	return input, nil
}

func validateIndexFilepath(filePath string) (string, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		f, err := os.Create(filePath)
		f.WriteString("# index\n[**shared files**](/share/)") // template index.md contents
		if err != nil {
			return DefaultConfig.IndexFilepath, err
		}
		defer f.Close()
	}

	return filePath, nil
}

func validateShareDirectoryFilepath(filePath string) (string, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) || os.IsPermission(err) {
		err = os.MkdirAll(filePath, 0755) // creates directories recursively with permissions set as 0755 (rwxr-xr-x)
		if err != nil {
			return DefaultConfig.ShareDirectoryFilepath, err
		}
	}

	if fileInfo, _ := os.Stat(filePath); fileInfo.IsDir() == false {
		return DefaultConfig.ShareDirectoryFilepath, fmt.Errorf("invalid share directory: %v is not a directory", filePath)
	}

	return filePath, nil
}
