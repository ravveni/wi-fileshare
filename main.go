package main

import (
	"log"
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

func renderIndex(w http.ResponseWriter, r *http.Request) {
	markdown, err := os.ReadFile("./index.md")
	if err != nil {
		log.Fatal(err)
	}

	html := blackfriday.MarkdownBasic(markdown)

	err = os.WriteFile("index.html", html, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	http.ServeFile(w, r, "./index.html")
}

func main() {
	fs := http.FileServer(http.Dir("./share"))

	http.HandleFunc("/", renderIndex)
	http.Handle("/share/", http.StripPrefix("/share/", fs))

	log.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
