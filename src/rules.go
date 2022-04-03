// Rules Utilities
package main

import (
	"fmt"
	"os"
)

func rulesOperation(args []string) {

	switch args[0] {

		case "reload":
			reloadRules()

		case "run":
			exitOnNil(args, 2)
			runRules(args[2])

		default:
			fmt.Printf(Red("Invalid argument: %s\n\n"), args[1])
			helpPrompt()
			os.Exit(0)
	}
}

func reloadRules() {
	unavailable()
}

func runRules(key string) {
	unavailable()
}
