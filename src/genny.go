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

	fmt.Println("\nLoading credentials config for " + Blue(user))

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
func loadProjects() {
	
	os.Chdir(GENNY_MAIN)

	projects := projects()

	for _, p := range projects {
		fmt.Println("Copying rules for " + Yellow(p))

		cmd := exec.Command("cp", "-rp", "../" + p + "/rules", "./rules/" + p + "/")
		cmd.Stderr = os.Stderr
		tail(cmd)
	}
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

// Start the Genny system.
func startGenny(parser Parser) {

	containers := parser.getFrom(1)

	os.Chdir(GENNY_MAIN)

	// load credentials and rules directories
	loadEnvironment("dev1")
	fmt.Println("")

	loadProjects()
	fmt.Println("")

	// create network and volume
	createDockerNetwork("mainproxy")
	createDockerVolume("mysql_data")

	fmt.Print("\nStarting Docker Containers...\n\n")

	if containers != nil {
		// start containers
		arguments := append([]string{"up", "-d"}, containers...)
		cmd := exec.Command("docker-compose", arguments...)
		cmd.Stderr = os.Stderr
		tail(cmd)
	} else {
		// start all containers
		cmd := exec.Command("docker-compose", "up", "-d")
		cmd.Stderr = os.Stderr
		tail(cmd)
	}

	// set back to current working dir
	os.Chdir(CURRENT_DIR)

	if parser.hasFlag("-f") {
		tailServiceLogs(parser)
	}
}

// Stop the Genny system.
func stopGenny(parser Parser) {

	containers := parser.getFrom(1)

	fmt.Print("\nStopping Docker Containers...\n\n")
	os.Chdir(GENNY_MAIN)

	if containers != nil {
		// stop containers
		arguments := append([]string{"stop"}, containers...)
		cmd := exec.Command("docker-compose", arguments...)
		cmd.Stderr = os.Stderr
		tail(cmd)

		fmt.Print("\nRemoving Docker Containers...\n\n")

		// remove containers
		arguments = append([]string{"rm", "-f"}, containers...)
		cmd = exec.Command("docker-compose", arguments...)
		cmd.Stderr = os.Stderr
		tail(cmd)
	} else {
		// stop all containers
		cmd := exec.Command("docker-compose", "stop")
		cmd.Stderr = os.Stderr
		tail(cmd)

		fmt.Print("\nRemoving Docker Containers...\n\n")

		// remove all containers
		cmd = exec.Command("docker-compose", "rm", "-f")
		cmd.Stderr = os.Stderr
		tail(cmd)
	}

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

	fmt.Print("\nTailing Service logs...\n\n")
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

	// docker log our containers
	arguments := append([]string{"logs", "-f"}, containersToLog...)
	cmd := exec.Command("docker-compose", arguments...)
	cmd.Stderr = os.Stderr
	tail(cmd)

	// set back to current working dir
	os.Chdir(CURRENT_DIR)
}
