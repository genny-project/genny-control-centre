package main

import (
	"fmt"
	"os"
	"os/exec"
)

var dependencies = []string{"qwandaq", "serviceq"}  
var services = []string{"bridge", "kogitoq2", "fyodor", "lauchy", "dropkick", "messages"}  

func repos() []string {

	var output []string
	output = append(output, dependencies...)
	output = append(output, services...)

	return output
}

func repoStatus() {

	// save current dir and find home
	current, _ := os.Getwd()
	home := os.Getenv("HOME")

	repos := repos()

	for i := 0; i < len(repos); i++ {

		// create path to rep
		repo := repos[i]
		path := home + "/projects/genny/" + repo

		// Move to repo and pull
		os.Chdir(path)
		branch, _ := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
		fmt.Printf("Project: %s \t Branch; %s\n", yellow(repo), green(string(branch)))

		cmd := exec.Command("git", "-c", "color.status=always", "status")
		tail(cmd)
		fmt.Println("")
	}

	// set back to current working dir
	os.Chdir(current)
}

func cloneRepos(version string) {

	// save current dir and find home
	current, _ := os.Getwd()
	home := os.Getenv("HOME")

	path := home + "/projects/genny"
	os.Chdir(path)

	repos := repos()

	for i := 0; i < len(repos); i++ {

		// create path to rep
		repo := repos[i]
		url := "git@github.com:genny-project/" + repo
		fmt.Println(yellow("Cloning " + repo + "..."))

		// clone repo
		cmd := exec.Command("git", "clone", "-b", version, url)
		tail(cmd)
		fmt.Println("")
	}

	// set back to current working dir
	os.Chdir(current)
}

func pullRepos() {

	// save current dir and find home
	current, _ := os.Getwd()
	home := os.Getenv("HOME")

	repos := repos()

	for i := 0; i < len(repos); i++ {

		// create path to rep
		repo := repos[i]
		path := home + "/projects/genny/" + repo
		fmt.Println(yellow("Pulling " + repo + "..."))

		// Move to repo and pull
		os.Chdir(path)
		cmd := exec.Command("git", "pull")
		tail(cmd)
		fmt.Println("")
	}

	// set back to current working dir
	os.Chdir(current)
}

func buildDockerImages() {

	// save current dir and find home
	current, _ := os.Getwd()
	home := os.Getenv("HOME")

	// build dependencies
	for i := 0; i < len(dependencies); i++ {

		// create path to repo
		repo := dependencies[i]
		path := home + "/projects/genny/" + repo
		fmt.Println(yellow("Building " + repo + "..."))

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
		path := home + "/projects/genny/" + repo
		fmt.Println(yellow("Building " + repo + "..."))

		// Move to repo and build image
		os.Chdir(path)
		cmd := exec.Command("./build-docker.sh")
		tail(cmd)
		fmt.Println("")
	}

	// set back to current working dir
	os.Chdir(current)
}

