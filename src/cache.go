// Cache Utilities
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Selector function for operations on the Genny cache
func cacheOperation(args []string) {

	switch args[0] {

		case "read":
			exitOnNil(args, 2)
			readCache(args[2])

		case "write":
			exitOnNil(args, 3)
			writeCache(args[2], args[3])

		case "remove":
			exitOnNil(args, 2)
			removeCache(args[2])

		default:
			fmt.Printf(Red("Invalid argument: %s\n\n"), args[1])
			helpPrompt()
			os.Exit(0)
	}
}

// Read the data stored in the cache for a given key.
func readCache(key string) {

	fmt.Printf("Reading %s from cache...\n", Yellow(key))

	token := getToken()

	uri := "http://localhost:4242/cache/" + key

	// create GET request
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}

	// set the request headers
	req.Header = http.Header {
		"Content-Type": []string{"application/json"},
		"Authorization": []string{"Bearer " + token},
	}

	// execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// read response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// close response and return token
	resp.Body.Close()

	output, err := PrettyString(string(responseBody))
	if err != nil {
		output = string(responseBody)
	}

	fmt.Println("\n" + output)
}

// Write a value to the cache for a given key.
func writeCache(key string, value string) {

	fmt.Printf("Writing value to cache for key %s...\n", Yellow(key))

	token := getToken()

	uri := "http://localhost:4242/cache/" + key

	// create POST request
	req, err := http.NewRequest("POST", uri, strings.NewReader(value))
	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}

	// set the request headers
	req.Header = http.Header {
		"Content-Type": []string{"application/json"},
		"Authorization": []string{"Bearer " + token},
	}

	// execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// read response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// close response and return token
	resp.Body.Close()

	if len(string(responseBody)) > 0 {
		fmt.Println("")
		fmt.Println(string(responseBody))
	}
}

// Remove an item stored in the cache using the item key.
func removeCache(key string) {

	fmt.Printf("Removing value in cache for key %s...\n", Yellow(key))

	token := getToken()

	uri := "http://localhost:4242/cache/" + key

	// create DELETE request
	req, err := http.NewRequest("DELETE", uri, nil)
	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}

	// set the request headers
	req.Header = http.Header {
		"Content-Type": []string{"application/json"},
		"Authorization": []string{"Bearer " + token},
	}

	// execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// read response
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// close response and return token
	resp.Body.Close()
}

