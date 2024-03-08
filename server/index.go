package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/russross/blackfriday/v2"
)

func (config WFConfig)renderIndex(w http.ResponseWriter, r *http.Request) {
	markdown, err := os.ReadFile(config.IndexFilepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	html := blackfriday.Run(markdown)

	err = os.WriteFile("index.html", html, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.ServeFile(w, r, "./index.html")
}
