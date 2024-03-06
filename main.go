package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/russross/blackfriday/v2"
)

const defaultPort = 8080

func renderIndex(w http.ResponseWriter, r *http.Request) {
	markdown, err := os.ReadFile("./index.md")
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

func startServer(port int) {
	serverPort := fmt.Sprint(":",port)
	fs := http.FileServer(http.Dir("./share"))

	http.HandleFunc("/", renderIndex)
	http.Handle("/share/", http.StripPrefix("/share/", fs))

	log.Println("Listening on ", serverPort)
	err := http.ListenAndServe(serverPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func validateInput(input string) (int, error) {
	num, err := strconv.Atoi(input)
	if err != nil {
        return -1, fmt.Errorf("Invalid number: %v", err)
	}

	if num < 0 || num > 65536 {
		return -1, fmt.Errorf("%d is not between 0 and 65536", num)
	}

   return num, nil
}

func main() {
	port := defaultPort
	args := os.Args[1:]
	if len(args) > 0 {
		validatedPort, err := validateInput(args[0])
		if err != nil {
			log.Fatal(err)
		}
		port = validatedPort
	}

	startServer(port)
}
