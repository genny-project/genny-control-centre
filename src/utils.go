package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// Helper funtion for unavailable commands.
func unavailable() {
	fmt.Println(Red("Sorry, this action is unavailable at this moment."))
	os.Exit(0)
}

// Generate a pretty json string from an input string.
// If the string cannot be parsed as json, the input 
// string will be returned with a non nil error.
func PrettyString(str string) (string, error) {

    var prettyJSON bytes.Buffer
    if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
        return str, err
    }

    return prettyJSON.String(), nil
}

// Generate a string of empty spaces with length.
func Spaces(length int) string {

	var output = ""
	for i := 0; i < length; i++ {
		output += " "
	}

	return output
}

// Generate a string of hyphens with length.
func Hyphens(length int) string {

	var output = ""
	for i := 0; i < length; i++ {
		output += "-"
	}

	return output
}

// Merge two maps into a single map. Values in map y will
// override any values in map x with the same key.
func Merge(x map[string]string, y map[string]string) map[string]string {

	output := make(map[string]string)

	for k, v := range y {
		output[k] = v
	}

	for k, v := range x {
		output[k] = v
	}

	return output
}

// Execute a command and tail the output logs.
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
