package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"net/http"
	"os"
)

type EntityAttribute struct {
	AttributeCode		string	`json:"attributeCode"`
	Created				string	`json:"created"`
	Updated				string	`json:"updated"`
	ValueBoolean		string	`json:"valueBoolean"`
	ValueInteger		string	`json:"valueInteger"`
	ValueLong			string	`json:"valueLong"`
	ValueDate			string	`json:"valueDate"`
	ValueDateTime		string	`json:"valueDateTime"`
	ValueString			string	`json:"valueString"`
}

type BaseEntity struct {
	Name					string	`json:"name"`
	Code					string	`json:"code"`
	Status					string	`json:"status"`
	BaseEntityAttributes	[]EntityAttribute 	`json:"baseEntityAttributes"`
}

func entityOperation(args []string) {

	switch args[0] {

		case "show":
			exitOnNil(args, 2)
			showEntity(args[2])

		case "watch":
			exitOnNil(args, 2)
			watchEntity(args[2])

		default:
			fmt.Printf(red("Invalid argument: %s\n\n"), args[1])
			help()
			os.Exit(0)
	}
}

func showEntity(code string) {
	fmt.Printf("Showing Entity %s...\n", yellow(code))

	token := getToken()

	uri := "http://localhost:4242/entity/" + code

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}

	req.Header = http.Header {
		"Content-Type": []string{"application/json"},
		"Authorization": []string{"Bearer " + token},
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// read response
	responseBody, err := ioutil.ReadAll(resp.Body)
	// _, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// close response and return token
	resp.Body.Close()

	entity := BaseEntity{}

	if err := json.Unmarshal(responseBody, &entity); err != nil {
		panic(err)
	}

	fmt.Println("\n Code - Name - Status")
	fmt.Println(entity.Code + " - " + entity.Name + " - " + entity.Status + "\n")

	lengths := make(map[string]int)

	for i := 0; i < len(entity.BaseEntityAttributes); i++ {

		ea := entity.BaseEntityAttributes[i]
		v := reflect.ValueOf(ea)
		typeOf := v.Type()

		for j := 0; j < typeOf.NumField(); j++ {

			field := typeOf.Field(j).Name
			value := fmt.Sprintf("%s", v.Field(j).Interface())
			length := len(value)

			// init as title length
			lengths[field] = len(field)

			if length > lengths[field] {
				lengths[field] = length
			}
		}
	}

	header := "|"
	output := ""

	for i := 0; i < len(entity.BaseEntityAttributes); i++ {

		ea := entity.BaseEntityAttributes[i]
		v := reflect.ValueOf(ea)
		typeOf := v.Type()
		line := "|"

		for j := 0; j < typeOf.NumField(); j++ {

			field := typeOf.Field(j).Name
			value := fmt.Sprintf("%s", v.Field(j).Interface())
			length := len(value)
			fmt.Println(length)

			if i == 0 {
				header += " "
				header += field
				header += Spaces(lengths[field] - len(field))
				header += " |"
			}

			line += " "
			line += value
			line += Spaces(lengths[field] - length)
			line += " |"
		}

		output += line

		if i < len(entity.BaseEntityAttributes) - 1 {
			output += "\n"
		}
	}

	divider := Hyphens(len(header))
	fmt.Println(divider)
	fmt.Println(header)
	fmt.Println(divider)
	fmt.Println(output)
	fmt.Println(divider)
}

func watchEntity(code string) {
	fmt.Printf("Watching entity %s...\n", yellow(code))
}

