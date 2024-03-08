package server

import (
	"fmt"
	"net/http"
)

func Start(config WFConfig) error {
	validatedConfig, err := validateConfig(config)
	if err != nil {
		return err
	}

	port := fmt.Sprint(":", validatedConfig.Port)
	fs := http.FileServer(http.Dir(validatedConfig.ShareDirectoryFilepath))

	http.HandleFunc("/", config.renderIndex)
	http.Handle("/share/", http.StripPrefix("/share/", fs))

	fmt.Println("listening on", port)
	return http.ListenAndServe(port, nil)
}
