// Token Utilities
package main

import (
	"encoding/json"
	"strings"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type KeycloakResponse struct {
	AccessToken 		string 		`json:"access_token"`
	ExpiresIn 			int 		`json:"expires_in"`
	RefreshExpiresIn 	int  		`json:"refresh_expires_in"`
	RefreshToken 		string 		`json:"refresh_token"`
	TokenType			string 		`json:"token_type"`
	NotBeforePolicy 	int 		`json:"not"-before-policy`
	SessionState 		string 		`json:"session_state"`
	Scope 				string 		`json:"scope"`
}

// Selector for token based operations.
func tokenOperation(args []string) {

	switch args[0] {

		case "get":
			token := getToken()
			fmt.Println("")
			fmt.Println(token)

		default:
			fmt.Printf(Red("Invalid argument: %s\n\n"), args[1])
			helpPrompt()
			os.Exit(0)
	}
}

// Get an access token from keycloak using the service 
// user environment variables.
func getToken() string {

	// grab keycloak access vars
	var realm string = os.Getenv("GENNY_REALM")
	var keycloakURL string = os.Getenv("GENNY_KEYCLOAK_URL")
	var clientID string = os.Getenv("GENNY_CLIENT_ID")
	var clientSecret string = os.Getenv("GENNY_CLIENT_SECRET")
	var serviceUsername string = os.Getenv("GENNY_SERVICE_USERNAME")
	var servicePassword string = os.Getenv("GENNY_SERVICE_PASSWORD")

	uri := keycloakURL + "/auth/realms/" + realm + "/protocol/openid-connect/token"

	// construct body of post request
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", serviceUsername)
	data.Set("password", servicePassword)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	// create POST request
	req, err := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}

	// set headers
	req.Header = http.Header {
		"Content-Type": []string{"application/x-www-form-urlencoded"},
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

	response := KeycloakResponse{}

	if err := json.Unmarshal(responseBody, &response); err != nil {
		panic(err)
	}

	return response.AccessToken
}
