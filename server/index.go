package server

import (
	"log"
	"net/http"
	"os"

	"github.com/russross/blackfriday/v2"
)

func (config WFConfig)renderIndex(w http.ResponseWriter, r *http.Request) {
	markdown, err := os.ReadFile(config.IndexFilepath)
	if err != nil {
		log.Fatal(err)
	}

	html := blackfriday.Run(markdown)
	tempHTMLPath := os.TempDir() + "index.html"

	err = os.WriteFile(tempHTMLPath, html, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	http.ServeFile(w, r, tempHTMLPath)
}
