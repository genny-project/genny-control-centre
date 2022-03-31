package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"github.com/joho/godotenv"
)

var dependencies = []string{"qwandaq", "serviceq"}  
var services = []string{"bridge", "kogitoq2", "fyodor", "lauchy", "dropkick", "messages"}  
var utilities = []string{"genny-main"}

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

func repos() []string {

	var output []string
	output = append(output, dependencies...)
	output = append(output, services...)
	output = append(output, utilities...)
	output = append(output, projects()...)

	return output
}

func loadCredentials(user string) {

	fmt.Println("\nLoading credentials config for " + Blue(user))

	credentialsPath := HOME + "/.genny/credentials/credentials-" + user

	err := godotenv.Load(credentialsPath + "/conf.env")

	if err != nil {
		fmt.Printf(Red("Could not load conf.env for %s, Err: %s"), user, err)
	}

	cmd := exec.Command("cp", credentialsPath + "/StoredCredential", GENNY_MAIN + "/google_credentials/StoredCredential")
	cmd.Stderr = os.Stderr
	tail(cmd)
}

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

func createDockerNetwork(network string) {

	cmd := exec.Command("docker", "network", "create", "--gateway", "172.18.0.1", "--subnet", "172.18.0.0/24", network)
	cmd.Stderr = os.Stderr
	cmd.Output()
	// tail(cmd)
}

func createDockerVolume(volume string) {

	cmd := exec.Command("docker", "volume", "create", volume)
	cmd.Stderr = os.Stderr
	tail(cmd)
}

func repoStatus() {

	repos := repos()

	for i := 0; i < len(repos); i++ {

		// create path to rep
		repo := repos[i]
		path := HOME + "/projects/genny/" + repo

		// Move to repo and pull
		os.Chdir(path)
		branch, _ := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
		fmt.Printf("Project: %s \t Branch; %s\n", Yellow(repo), Green(string(branch)))

		cmd := exec.Command("git", "-c", "color.status=always", "status")
		tail(cmd)
		fmt.Println("")
	}

	// set back to current working dir
	os.Chdir(CURREND_DIR)
}

func cloneRepos(version string) {

	path := HOME + "/projects/genny"
	os.Chdir(path)

	repos := repos()

	for i := 0; i < len(repos); i++ {

		// create path to rep
		repo := repos[i]
		url := "git@github.com:genny-project/" + repo
		fmt.Println(Yellow("Cloning " + repo + "..."))

		// clone repo
		cmd := exec.Command("git", "clone", "-b", version, url)
		tail(cmd)
		fmt.Println("")
	}

	// set back to current working dir
	os.Chdir(CURREND_DIR)
}

func pullRepos() {

	repos := repos()

	for i := 0; i < len(repos); i++ {

		// create path to rep
		repo := repos[i]
		path := HOME + "/projects/genny/" + repo
		fmt.Println(Yellow("Pulling " + repo + "..."))

		// Move to repo and pull
		os.Chdir(path)
		cmd := exec.Command("git", "pull")
		tail(cmd)
		fmt.Println("")
	}

	// set back to current working dir
	os.Chdir(CURREND_DIR)
}

func checkoutRepos(branch string) {

	repos := repos()

	for i := 0; i < len(repos); i++ {

		// create path to rep
		repo := repos[i]
		path := HOME + "/projects/genny/" + repo
		fmt.Println(Yellow("Checking out " + repo + "to " + branch + "..."))

		os.Chdir(path)
		// stash
		cmd := exec.Command("git", "stash")
		tail(cmd)
		// checkout
		cmd = exec.Command("git", "checkout", branch)
		tail(cmd)
		fmt.Println("")
	}

	// set back to current working dir
	os.Chdir(CURREND_DIR)
}

func buildDockerImages() {

	// build dependencies
	for i := 0; i < len(dependencies); i++ {

		// create path to repo
		repo := dependencies[i]
		path := HOME + "/projects/genny/" + repo
		fmt.Println(Yellow("Building " + repo + "..."))

		// Move to repo and build
		os.Chdir(path)
		cmd := exec.Command("./build.sh")
		tail(cmd)
		fmt.Println("")
	}

	// build services
	for i := 0; i < len(services); i++ {

		// create path to repo
		repo := services[i]
		path := HOME + "/projects/genny/" + repo
		fmt.Println(Yellow("Building " + repo + "..."))

		// Move to repo and build image
		os.Chdir(path)
		cmd := exec.Command("./build-docker.sh")
		tail(cmd)
		fmt.Println("")
	}

	// set back to current working dir
	os.Chdir(CURREND_DIR)
}

func startGenny(containers []string) {

	os.Chdir(GENNY_MAIN)

	// load credentials and rules directories
	loadCredentials("dev1")
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
	os.Chdir(CURREND_DIR)
}

func stopGenny(containers []string) {

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
	os.Chdir(CURREND_DIR)
}

func restartGenny(containers []string) {

	// stop and start containers
	stopGenny(containers)
	startGenny(containers)
}

