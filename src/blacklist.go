package main

import (
	"fmt"
	"os"
)

func blacklistOperation(args []string) {

	switch args[0] {

		case "delete":
			exitOnNil(args, 2)
			deleteBlacklist(args[2])

		default:
			fmt.Printf(red("Invalid argument: %s\n\n"), args[1])
			helpPrompt()
			os.Exit(0)
	}
}

func deleteBlacklist(id string) {
	fmt.Printf("Reloading Rules...\n")
}

