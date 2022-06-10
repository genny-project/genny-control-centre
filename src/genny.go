// Genny System Utilities
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"github.com/joho/godotenv"
)

var oldDependencies = []string{"qwanda", "qwanda-utils", "bootxport", "genny-verticle-rules", "genny-rules", "qwanda-services"}
var oldServices = []string{"wildfly-qwanda-service", "wildfly-rulesservice"}
var services = []string{"gennyq", "alyson"}
var utilities = []string{"genny-main"}

// Return a list of all project repositories
func projects() []string {

   files, _ := ioutil.ReadDir("../")

   var output []string
   for _, f := range files {
	   if strings.HasPrefix(f.Name(), "prj_") {
		   output = append(output, f.Name())
	   }
   }

	return output
}

// Return a slice of all repository names
func repos() []string {

	var output []string
	output = append(output, oldDependencies...)
	output = append(output, oldServices...)
	output = append(output, services...)
	output = append(output, utilities...)
	output = append(output, projects()...)

	return output
}

// Load a user specific environment.
func loadEnvironment(user string) {

	fmt.Println("Loading credentials config for " + Blue(user))

	var envMap map[string]string
	var credMap map[string]string

	envMap, err := godotenv.Read(ENV_FILE)
	if err != nil {
		fmt.Printf(Red("Could not read %s, Err: %s"), ENV_FILE, err)
	}

	credentialsPath := HOME + "/.genny/credentials/credentials-" + user

    credMap, err = godotenv.Read(credentialsPath + "/conf.env")
	if err != nil {
		fmt.Printf(Red("Could not load conf.env for %s, Err: %s"), user, err)
	}

	output := Merge(envMap, credMap)

	err = godotenv.Write(output, GENNY_MAIN + "/genny.env")
	if err != nil {
		panic(err)
	}
}

// Load project rules into the central rules directory.
func loadProjectRules() {
	
	os.Chdir(GENNY_MAIN)

	projects := projects()

	for _, p := range projects {
		fmt.Println("Copying rules for " + Yellow(p))

		cmd := exec.Command("cp", "-rp", "../" + p + "/rules", "./rules/" + p + "/")
		cmd.Stderr = os.Stderr
		tail(cmd)
	}
}

func loadProtobufs() {

	os.Chdir(GENNY_MAIN)

	PERSISTENCE_FOLDER := GENNY_MAIN + "/target/protobuf"
	TARGET_PROTOS := GENNY_HOME + "/gennyq/kogitoq/gadaq/target/classes/META-INF/resources/persistence/protobuf/*.proto"

	fmt.Println("Copying protobufs for " + Yellow("gadaq"))

	cmd := exec.Command("/bin/cp", "-f", TARGET_PROTOS, PERSISTENCE_FOLDER)
	cmd.Stderr = os.Stderr
	tail(cmd)
}

// Create a docker network
func createDockerNetwork(network string) {

	cmd := exec.Command("docker", "network", "create", "--gateway", "172.18.0.1", "--subnet", "172.18.0.0/24", network)
	cmd.Stderr = os.Stderr
	cmd.Output()
}

// Create a docker volume
func createDockerVolume(volume string) {

	cmd := exec.Command("docker", "volume", "create", volume)
	cmd.Stderr = os.Stderr
	tail(cmd)
}

// Perform a git status on all repositories.
func repoStatus() {

	repos := repos()

	for i := 0; i < len(repos); i++ {

		// create path to rep
		repo := repos[i]
		path := HOME + "/projects/genny/" + repo

		// Move to repo and get branch
		os.Chdir(path)
		branch, _ := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
		fmt.Printf("Project: %s \t Branch; %s\n", Yellow(repo), Green(string(branch)))

		cmd := exec.Command("git", "-c", "color.ui=always", "status")
		tail(cmd)
		fmt.Println("")
	}

	// set back to current working dir
	os.Chdir(CURRENT_DIR)
}

// Perform a git clone on all repositories.
func cloneRepos(parser Parser) {

	version := ""
	if len(parser.coreArgs) > 1 {
		version = parser.get(1)
	}

	path := HOME + "/projects/genny"
	os.Chdir(path)

	repos := repos()

	for i := 0; i < len(repos); i++ {

		// create path to rep
		repo := repos[i]
		url := "git@github.com:genny-project/" + repo
		fmt.Println(Yellow("Cloning " + repo + "..."))

		// clone repo
		cmd := exec.Command("git", "-c", "color.ui=always", "clone", "-b", version, url)
		tail(cmd)
		fmt.Println("")
	}

	// set back to current working dir
	os.Chdir(CURRENT_DIR)
}

// Perform a git pull on all repositories.
func pullRepos() {

	repos := repos()

	for i := 0; i < len(repos); i++ {

		// create path to rep
		repo := repos[i]
		path := HOME + "/projects/genny/" + repo
		os.Chdir(path)

		fmt.Println(Yellow("Pulling " + repo + "..."))

		// perform a pull
		cmd := exec.Command("git", "-c", "color.ui=always", "pull")
		tail(cmd)
		fmt.Println("")
	}

	// set back to current working dir
	os.Chdir(CURRENT_DIR)
}

// Perform a git checkout on all repositories.
func checkoutRepos(parser Parser) {

	branch := parser.get(0)
	repos := repos()

	for i := 0; i < len(repos); i++ {

		// create path to rep
		repo := repos[i]
		path := HOME + "/projects/genny/" + repo
		fmt.Println(Yellow("Checking out " + repo + " to " + branch + "..."))

		os.Chdir(path)
		// stash
		cmd := exec.Command("git", "stash")
		tail(cmd)
		// checkout
		cmd = exec.Command("git", "-c", "color.ui=always", "checkout", branch)
		tail(cmd)
		fmt.Println("")
	}

	// set back to current working dir
	os.Chdir(CURRENT_DIR)
}

// Build all required Genny docker images.
func buildDockerImages() {

	// TODO: add more repos to this function

	// build dependencies
	for i := 0; i < len(oldDependencies); i++ {

		// create path to repo
		repo := oldDependencies[i]
		path := GENNY_HOME + "/" + repo
		fmt.Println(Yellow("Building " + repo + "..."))

		// Move to repo and build
		os.Chdir(path)
		cmd := exec.Command("./build.sh")
		tail(cmd)
		fmt.Println("")
	}

	for i := 0; i < len(oldServices); i++ {

		// create path to repo
		repo := oldServices[i]
		path := GENNY_HOME + "/" + repo
		fmt.Println(Yellow("Building " + repo + "..."))

		// Move to repo and build
		os.Chdir(path)
		cmd := exec.Command("./build-docker.sh")
		tail(cmd)
		fmt.Println("")
	}

	// build services
	for i := 0; i < len(services); i++ {

		// create path to repo
		repo := services[i]
		path := GENNY_HOME + "/" + repo

		// handle alyson path
		if repo == "alyson" {
			path = path + "/scripts"
		}

		fmt.Println(Yellow("Building " + repo + "..."))

		// Move to repo and build image
		os.Chdir(path)
		cmd := exec.Command("./build-docker.sh")
		tail(cmd)
		fmt.Println("")
	}

	// set back to current working dir
	os.Chdir(CURRENT_DIR)
}

// Load project data using bootq
func loadProjectData(parser Parser) {

	projects := parser.getFrom(1)

	os.Chdir(GENNY_MAIN)

	// load credentials, rules directories and protos
	Banner("Loading project data...")

	if projects != nil {

	}

	// set back to current working dir
	os.Chdir(CURRENT_DIR)
}

// Get all the compose file arguments in use
func getComposeFileArguments() []string {

	PRODUCT_CODES := os.Getenv("PRODUCT_CODES")
	products := strings.Split(PRODUCT_CODES, ",")

	arguments := []string{"-f", "docker-compose.yml"}

	// add any product compose files
	for _, product := range products {
		file := GENNY_HOME + "/products/" + product + "/docker-compose.yml"
		arguments = append(arguments, []string{"-f", file}...)
	}

	return arguments
}

// Start the Genny system.
func startGenny(parser Parser) {

	containers := parser.getFrom(1)

	os.Chdir(GENNY_MAIN)

	// load credentials, rules directories and protos
	Banner("Loading System Environment...")

	loadEnvironment("dev1")
	loadProjectRules()
	// loadProtobufs()

	// create network and volume
	createDockerNetwork("mainproxy")
	createDockerVolume("mysql_data")

	Banner("Starting Docker Containers...")

	// create compose arguments
	arguments := getComposeFileArguments();
	arguments = append(arguments, []string{"up", "-d"}...)

	// select specific containers
	if containers != nil {
		arguments = append(arguments, containers...)
	}

	fmt.Println("docker-compose " + strings.Join(arguments[:], " "))

	// run command with args
	cmd := exec.Command("docker-compose", arguments...)
	cmd.Stderr = os.Stderr
	tail(cmd)

	// set back to current working dir
	os.Chdir(CURRENT_DIR)

	if parser.hasFlag("-f") {
		tailServiceLogs(parser)
	}
}

// Stop the Genny system.
func stopGenny(parser Parser) {

	containers := parser.getFrom(1)

	os.Chdir(GENNY_MAIN)

	arguments := getComposeFileArguments();
	stop := append(arguments, []string{"stop"}...)
	rm := append(arguments, []string{"rm", "-f"}...)

	if containers != nil {
		stop = append(stop, containers...)
		rm = append(rm, containers...)
	}

	Banner("Stopping Docker Containers...")
	cmd := exec.Command("docker-compose", stop...)
	cmd.Stderr = os.Stderr
	tail(cmd)

	Banner("Removing Docker Containers...")
	cmd = exec.Command("docker-compose", rm...)
	cmd.Stderr = os.Stderr
	tail(cmd)

	// set back to current working dir
	os.Chdir(CURRENT_DIR)
}

// Restart the Genny system.
func restartGenny(parser Parser) {

	// stop and start containers
	stopGenny(parser)
	startGenny(parser)
}

// Watch the logs for a set of Genny containers.
func tailServiceLogs(parser Parser) {

	containers := parser.getFrom(1)

	Banner("Tailing Service logs...")
	os.Chdir(GENNY_MAIN)

	// get all active containers
	out, err := exec.Command("docker", "ps", "-a", "--format", "{{.Names}}").Output()
	if err != nil {
		panic(err)
	}

	// convert from string to array
	activeContainers := string(out)
	activeContainerArray := strings.Split(activeContainers, "\n")

	var containersToLog []string

	// diy grepper, check if container has substring
	for _, a := range activeContainerArray {
		add := false
		for _, c := range containers {
			if strings.Contains(a, c) {
				add = true
			}
		}

		if add {
			containersToLog = append(containersToLog, a)
		}
	}

	if len(containersToLog) == 0 {
		fmt.Println("No containers to log!")
	} else {
		// docker log our containers
		arguments := getComposeFileArguments();
		arguments = append(arguments, []string{"logs", "-f"}...)
		arguments = append(arguments, containersToLog...)
		cmd := exec.Command("docker-compose", arguments...)
		cmd.Stderr = os.Stderr
		tail(cmd)
	}

	// set back to current working dir
	os.Chdir(CURRENT_DIR)
}
