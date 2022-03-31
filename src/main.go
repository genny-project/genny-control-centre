package main

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
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

var HOME string
var GENNY_HOME string
var GENNY_MAIN string
var ENV_FILE string
var CURREND_DIR string

func main() {

	// genny location vars
	HOME = os.Getenv("HOME")
	GENNY_HOME = os.Getenv("GENNY_HOME")
	GENNY_MAIN = os.Getenv("GENNY_MAIN")
	ENV_FILE = os.Getenv("ENV_FILE")

	CURREND_DIR, _ = os.Getwd()

	err := godotenv.Load(ENV_FILE)
	if err != nil {
		fmt.Printf(Red("Could not load %s, Err: %s"), ENV_FILE, err)
	}

	var version string = "1.0.0"

	// grab arguments
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please specify valid arguments!")
		os.Exit(0)
	}

	// single length commands
	switch args[0] {

	case "version":
		fmt.Printf("%s\n", version)
		os.Exit(0)

	case "help":
		help()
		os.Exit(0)

	case "status":
		repoStatus()
		os.Exit(0)

	case "clone":
		version := ""
		if len(args) > 1 {
			version = args[1]
		}
		cloneRepos(version)
		os.Exit(0)

	case "pull":
		pullRepos()
		os.Exit(0)

	case "checkout":
		if len(args) > 1 {
			checkoutRepos(args[1])
		} else {
			fmt.Println(Red("Please provide a valid branch to checkout!"))
		}
		os.Exit(0)

	case "build":
		buildDockerImages()
		os.Exit(0)

	case "start":
		if len(args) > 1 {
			startGenny(args[1:])
		} else {
			startGenny(nil)
		}
		os.Exit(0)

	case "stop":
		if len(args) > 1 {
			stopGenny(args[1:])
		} else {
			stopGenny(nil)
		}
		os.Exit(0)

	case "restart":
		if len(args) > 1 {
			restartGenny(args[1:])
		} else {
			restartGenny(nil)
		}
		os.Exit(0)

	case "logs":
		if len(args) > 1 {
			tailServiceLogs(args[1:])
		} else {
			tailServiceLogs(nil)
		}
		os.Exit(0)
	}

	// multi length commands
	if len(args) > 1 {

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
			fmt.Printf(Red("Invalid argument: %s\n"), args[1])
			helpPrompt()
		}
	}

	// finish
	os.Exit(0)
}

func helpPrompt() {

	fmt.Println("\nTo see a list of valid commands, run: \n    gctl help")
}

func help() {

	fmt.Println("")
	fmt.Println(Yellow("Genny System Control Centre"))
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("    gctl <operation> <command> <data>")
	fmt.Println("    gctl <command> <data>")
	fmt.Println("    gctl <command>")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("    gctl read cache SBE_USERS")
	fmt.Println("    gctl restart bridge")
	fmt.Println("    gctl status")
	fmt.Println("")
	fmt.Println("")
	fmt.Println(Blue("Commands: "))
	fmt.Println("    help         Show valid commands")
	fmt.Println("    version      Print the version")
	fmt.Println("")
	fmt.Println("    status       Show Git status of Genny repositories")
	fmt.Println("    clone        Git Clone Genny repositories")
	fmt.Println("    pull         Git Pull Genny repositories")
	fmt.Println("")
	fmt.Println("    build        Build Genny docker containers")
	fmt.Println("    start        Start docker containers")
	fmt.Println("    stop         Stop docker containers")
	fmt.Println("    restart      Restart docker containers")
	fmt.Println("")
	fmt.Println("    cache        Perfrom a cache operation")
	fmt.Println("    entity       Perform an entity operation")
	fmt.Println("    search       Perform a search operation")
	fmt.Println("    token        Perform a token operation")
	fmt.Println("    rules        Perform a rules operation")
	fmt.Println("    blacklist    Perform a blacklist operation")
	fmt.Println("")
	fmt.Println("")
	fmt.Println(Blue("Cache Operations: "))
	fmt.Println("    read         Read an item from the cache")
	fmt.Println("    write        Write json to the cache")
	fmt.Println("    remove       Remove an item from the cache")
	fmt.Println("")
	fmt.Println(Blue("Entity Operations: "))
	fmt.Println("    show         Show the state of an entity in the database")
	fmt.Println("    watch        Watch the state of an entity in the database")
	fmt.Println("")
	fmt.Println(Blue("Search Operations: "))
	fmt.Println("    fetch        Fetch entities using a Genny search")
	fmt.Println("    count        Count entities using a Genny search")
	fmt.Println("")
	fmt.Println(Blue("Token Operations: "))
	fmt.Println("    get          Get an access token")
	fmt.Println("")
	fmt.Println(Blue("Rules Operations: "))
	fmt.Println("    reload       Reload the rules engine")
	fmt.Println("    run          Run a rule group")
	fmt.Println("")
	fmt.Println(Blue("Blacklist Operations: "))
	fmt.Println("    delete       Delete a blacklist by user id")
	fmt.Println("")
}

func exitOnNil(args []string, index int) {

	if len(args) <= index {
		fmt.Println(Red("Incorrect number of arguments!"))
		helpPrompt()
		os.Exit(0)
	}
}

