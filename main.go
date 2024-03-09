package main

import (
	"log"
	"runtime"

	"github.com/ravveni/wi-fileshare/server"
)

func main() {
	var config server.WFConfig

	if runtime.GOOS == "windows" {
		config = server.DefaultWindowsConfig
	} else {
		config = server.DefaultUnixConfig
	}

	log.Fatal(server.Start(config))
}
