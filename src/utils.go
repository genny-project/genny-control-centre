package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func unavailable() {
	fmt.Println(Red("Sorry, this action is unavailable at this moment."))
	os.Exit(0)
}

func PrettyString(str string) (string, error) {

    var prettyJSON bytes.Buffer
    if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
        return "", err
    }

    return prettyJSON.String(), nil
}

func Spaces(length int) string {

	var output = ""
	for i := 0; i < length; i++ {
		output += " "
	}

	return output
}

func Hyphens(length int) string {

	var output = ""
	for i := 0; i < length; i++ {
		output += "-"
	}

	return output
}

func tail(cmd *exec.Cmd) {

	// create a pipe for the output of the script
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe for Cmd: " + err.Error())
		panic(err)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting Cmd: " + err.Error())
		panic(err)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for Cmd: " + err.Error())
		panic(err)
	}
}
