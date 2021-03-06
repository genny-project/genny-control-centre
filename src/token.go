// Token Utilities
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type KeycloakResponse struct {
	AccessToken 		string 		`json:"access_token"`
	ExpiresIn 			int 		`json:"expires_in"`
	RefreshExpiresIn 	int  		`json:"refresh_expires_in"`
	RefreshToken 		string 		`json:"refresh_token"`
	TokenType			string 		`json:"token_type"`
	NotBeforePolicy 	int 		`json:"not-before-policy"`
	SessionState 		string 		`json:"session_state"`
	Scope 				string 		`json:"scope"`
}

// Get an access token from keycloak using the service 
// user environment variables.
func getToken(user string) string {

	// grab keycloak access vars
	var realm string = os.Getenv("GENNY_KEYCLOAK_REALM")
	var keycloakURL string = os.Getenv("GENNY_KEYCLOAK_URL")
	var username string;
	var password string;
	var clientID string;
	var clientSecret string = "";


	if user == "service" {

		clientID = os.Getenv("GENNY_CLIENT_ID")
		clientSecret = os.Getenv("GENNY_CLIENT_SECRET")
		username = os.Getenv("GENNY_SERVICE_USERNAME")
		password = os.Getenv("GENNY_SERVICE_PASSWORD")

	} else if user == "test" {

		clientID = os.Getenv("GENNY_TEST_CLIENT_ID")
		username = os.Getenv("GENNY_TEST_USERNAME")
		password = os.Getenv("GENNY_TEST_PASSWORD")
	}

	uri := keycloakURL + "/auth/realms/" + realm + "/protocol/openid-connect/token"

	// construct body of post request
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", username)
	data.Set("password", password)
	data.Set("client_id", clientID)

	if clientSecret != "" {
		data.Set("client_secret", clientSecret)
	}

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
