package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ravveni/wi-fileshare/server"
)

func main() {
	config := server.DefaultConfig

	args := os.Args[1:]
	if len(args) > 0 {
		if args[0] == "--help" || args[0] == "-h" || args[0] == "man" {
			printUsageGuide()
			return
		}

		for _, v := range args {
			arg := strings.SplitN(v, "=", 2)
			if arg[0] == "-p" {
				config.Port = arg[1]
			} else if arg[0] == "-i" {
				config.IndexFilepath = arg[1]
			} else if arg[0] == "-s" {
				config.ShareDirectoryFilepath = arg[1]
			} else {
				fmt.Println("Invalid action, 'wi-fileshare --help' for usage guide.")
				return
			}
		}
	}

	fmt.Print(server.Start(config))
}

func printUsageGuide() {
	fmt.Println("           _       _____ __          __")
	fmt.Println(" _      __(_)     / __(_) /__  _____/ /_  ____ _________")
	fmt.Println("| | /| / / /_____/ /_/ / / _ \\/ ___/ __ \\/ __ `/ ___/ _ \\")
	fmt.Println("| |/ |/ / /_____/ __/ / /  __(__  ) / / / /_/ / /  /  __/")
	fmt.Println("|__/|__/_/     /_/ /_/_/\\___/____/_/ /_/\\__,_/_/   \\___/")
	fmt.Println("")
	fmt.Println("# basic:")
	fmt.Println("$ ./wi-fileshare")
	fmt.Println("")
	fmt.Println("within the current working directory:")
	fmt.Println("- creates an `index.md` for landing page, skips if exists")
	fmt.Println("- creates a `share/` directory for filesharing, skips if exists")
	fmt.Println("- converts `index.md` to `index.html`")
	fmt.Println("- serves on port `8080`")
	fmt.Println("")
	fmt.Println("# custom:")
	fmt.Println("$ ./wi-fileshare -p=PORT -i=INDEX.MD-FILEPATH -s=SHARE-DIRECTORY-FILEPATH")
	fmt.Println("")
	fmt.Println("all arguments are optional, defaults same from basic usage")
}
