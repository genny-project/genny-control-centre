package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/joho/godotenv"
	"github.com/fatih/color"
)

// gctl version
var version string = "1.2.1"

// environment variables
var HOME string
var GENNY_HOME string
var GENNY_MAIN string
var ENV_FILE string
var CURRENT_DIR string
var CONTAINER_ENGINE string

var red = color.New(color.FgRed).PrintfFunc()
var blue = color.New(color.FgBlue).PrintfFunc()
var yellow = color.New(color.FgYellow).PrintfFunc()

// Main function execution.
func main() {

	// genny location vars
	HOME = os.Getenv("HOME")
	GENNY_HOME = os.Getenv("GENNY_HOME")
	GENNY_MAIN = os.Getenv("GENNY_MAIN")
	ENV_FILE = os.Getenv("GENNY_ENV_FILE")
	CONTAINER_ENGINE = os.Getenv("CONTAINER_ENGINE")

	CURRENT_DIR, _ = os.Getwd()

	err := godotenv.Load(ENV_FILE)
	if err != nil {
		fmt.Printf(Red("Could not load %s, Err: %s"), ENV_FILE, err)
	}

	// grab arguments
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please specify valid arguments!")
		os.Exit(0)
	}

	// parse command line arguments and flags
	parser := Parser{}
	parser.parse(args)

	if parser.containsOne("version") {

		fmt.Printf("%s\n", version)

	} else if parser.containsOne("help") {

		help()

	} else if parser.containsOne("status") {

		repoStatus()

	} else if parser.containsOne("clone") {

		cloneRepos(parser)

	} else if parser.containsOne("pull") {

		pullRepos()

	} else if parser.containsOne("checkout") {

		checkoutRepos(parser)

	} else if parser.containsOne("build") {

		buildDockerImages()

	} else if parser.containsOne("load") {

		loadProjectData(parser)

	} else if parser.containsOne("start") {

		startGenny(parser)

	} else if parser.containsOne("stop") {

		stopGenny(parser)

	} else if parser.containsOne("restart") {

		restartGenny(parser)

	} else if parser.containsOne("logs") {

		tailServiceLogs(parser)

	// cache operations
	} else if parser.containsTwo("read", "cache") {

		readCache(parser.get(2))

	} else if parser.containsTwo("write", "cache") {

		writeCache(parser.get(2), parser.get(3), parser.get(4))

	} else if parser.containsTwo("remove", "cache") {

		removeCache(parser.get(2))

	} else if parser.containsTwo("show", "entity") {

		showEntity(parser.get(2))

	} else if parser.containsTwo("watch", "entity") {

		watchEntity(parser.get(2))

	} else if parser.containsTwo("fetch", "search") {

		fetch(parser.get(2))

	} else if parser.containsTwo("count", "search") {

		count(parser.get(2))

	} else if parser.containsTwo("get", "token") {

		token := getToken(parser.get(2))
		fmt.Println(token)

	} else if parser.containsTwo("reload", "rules") {

		reloadRules()

	} else if parser.containsTwo("run", "rules") {

		runRules(parser.get(2))

	} else if parser.containsTwo("delete", "blacklist") {

		deleteBlacklist(parser.get(2))

	} else {

		red("Invalid command: %s\n", strings.Join(parser.getFrom(0), " "))
		helpPrompt()

	}

	// finish
	os.Exit(0)
}

// Print a prompt for finding the help command.
func helpPrompt() {

	fmt.Println("\nTo see a list of valid commands, run: \n    gctl help")
}

// Print a help description.
func help() {

	fmt.Println("")
	yellow("Genny System Control Centre\n")
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
	blue("Commands:\n")
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
	blue("Cache Operations:\n")
	fmt.Println("    read         Read an item from the cache")
	fmt.Println("    write        Write json to the cache")
	fmt.Println("    remove       Remove an item from the cache")
	fmt.Println("")
	blue("Entity Operations:\n")
	fmt.Println("    show         Show the state of an entity in the database")
	fmt.Println("    watch        Watch the state of an entity in the database")
	fmt.Println("")
	blue("Search Operations:\n")
	fmt.Println("    fetch        Fetch entities using a Genny search")
	fmt.Println("    count        Count entities using a Genny search")
	fmt.Println("")
	blue("Token Operations:\n")
	fmt.Println("    get          Get an access token")
	fmt.Println("")
	blue("Rules Operations:\n")
	fmt.Println("    reload       Reload the rules engine")
	fmt.Println("    run          Run a rule group")
	fmt.Println("")
	blue("Blacklist Operations:\n")
	fmt.Println("    delete       Delete a blacklist by user id")
	fmt.Println("")
}

// Exit the program if the length of arguments is not adequate.
func exitOnNil(args []string, index int) {

	if len(args) <= index {
		fmt.Println(Red("Incorrect number of arguments!"))
		helpPrompt()
		os.Exit(0)
	}
}

