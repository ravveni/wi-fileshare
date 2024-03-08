package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/russross/blackfriday/v2"
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

	log.Println("Listening on ", port)
	return http.ListenAndServe(port, nil)
}

func (config WFConfig)renderIndex(w http.ResponseWriter, r *http.Request) {
	markdown, err := os.ReadFile(config.IndexFilepath)
	if err != nil {
		log.Fatal(err)
	}

	html := blackfriday.Run(markdown)

	err = os.WriteFile("index.html", html, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	http.ServeFile(w, r, "./index.html")
}
