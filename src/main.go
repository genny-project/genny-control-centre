package main

import (
	"os"
	"fmt"
	// "github.com/joho/godotenv"
)

// environment variables
var realm string
var keycloakURL string
var clientID string
var clientSecret string
var serviceUsername string
var servicePassword string
var infinispanURL string
var infinispanUsername string
var infinispanPassword string

func main() {

	// err := godotenv.Load("~/projects/genny/genny-main/genny.env")
	// if err != nil {
	// 	fmt.Printf(red("Could not load genny.env. Err: %s"), err)
	// }

	var version string = "1.0.0"

	// grab arguments
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please specify valid arguments!")
		os.Exit(0)
	}

	// single length commands
	if len(args) == 1 {

		switch args[0] {

			case "version":
				fmt.Printf("%s\n", version)

			case "help":
				help()

			case "status":
				repoStatus()

			case "clone":
				version := ""
				if len(args) > 1 {
					version = args[1]
				}
				cloneRepos(version)

			case "pull":
				pullRepos()

			case "build":
				buildDockerImages()

			default:
				fmt.Println("Unknown command: " + args[0])
				helpPrompt()
		}

		os.Exit(0)
	}

	// multi length commands
	switch args[1] {

		case "cache":
			exitOnNil(args, 1)
			cacheOperation(args)

		case "entity":
			exitOnNil(args, 1)
			entityOperation(args)

		case "search":
			exitOnNil(args, 1)
			searchOperation(args)

		case "token":
			exitOnNil(args, 1)
			tokenOperation(args)

		case "rules":
			exitOnNil(args, 1)
			rulesOperation(args)

		case "blacklist":
			exitOnNil(args, 1)
			blacklistOperation(args)

		default:
			fmt.Printf(red("Invalid argument: %s\n"), args[1])
			helpPrompt()
			os.Exit(0)
	}
}

func helpPrompt() {

	fmt.Println("\nTo see a list of valid commands, run: \n    gctl help")
}

func help() {

	fmt.Println("Genny System Control Centre")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Usage: gctl <operation> <command> <data>")
	fmt.Println("")
	fmt.Println("Example: gctl read cache SBE_USERS")
	fmt.Println("")
	fmt.Println("")
	fmt.Println(blue("Commands: "))
	fmt.Println("    help         Show valid commands")
	fmt.Println("    version      Print the version")
	fmt.Println("    cache        Perfrom a cache operation")
	fmt.Println("    entity       Perform an entity operation")
	fmt.Println("    search       Perform a search operation")
	fmt.Println("    token        Perform a token operation")
	fmt.Println("    rules        Perform a rules operation")
	fmt.Println("    blacklist    Perform a blacklist operation")
	fmt.Println("")
	fmt.Println("")
	fmt.Println(blue("Cache Operations: "))
	fmt.Println("    read         Read an item from the cache")
	fmt.Println("    write        Write json to the cache")
	fmt.Println("    remove       Remove an item from the cache")
	fmt.Println("")
	fmt.Println(blue("Entity Operations: "))
	fmt.Println("    show         Show the state of an entity in the database")
	fmt.Println("    watch        Watch the state of an entity in the database")
	fmt.Println("")
	fmt.Println(blue("Search Operations: "))
	fmt.Println("    fetch        Fetch entities using a Genny search")
	fmt.Println("    count        Count entities using a Genny search")
	fmt.Println("")
	fmt.Println(blue("Token Operations: "))
	fmt.Println("    get          Get an access token")
	fmt.Println("")
	fmt.Println(blue("Rules Operations: "))
	fmt.Println("    reload       Reload the rules engine")
	fmt.Println("    run          Run a rule group")
	fmt.Println("")
	fmt.Println(blue("Blacklist Operations: "))
	fmt.Println("    delete       Delete a blacklist by user id")
	fmt.Println("")
}

func exitOnNil(args []string, index int) {

	if len(args) <= index {
		fmt.Println(red("Incorrect number of arguments!"))
		helpPrompt()
		os.Exit(0)
	}
}

