package main

import (
	"fmt"
	"os"
)

func searchOperation(args []string) {

	switch args[0] {

		case "fetch":
			exitOnNil(args, 3)
			fetch(args[2])

		case "count":
			exitOnNil(args, 3)
			count(args[2])

		default:
			fmt.Printf(Red("Invalid argument: %s\n\n"), args[1])
			helpPrompt()
			os.Exit(0)
	}
}

func fetch(json string) {
	unavailable()
}

func count(json string) {
	unavailable()
}
