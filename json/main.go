package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //alias to send to backend
	price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON")
	// EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {
	courses := []course{
		{"Reactjs", 300, "lco", "abc123", []string{"web-dev", "js"}},
		{"MERN", 400, "lco", "mer123", []string{"web-dev", "js"}},
		{"RN", 500, "lco", "rn123", nil},
	}

	finalJSON, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJSON))
}

func DecodeJSON() {
	jsonData := []byte(`
	{
                "coursename": "Reactjs",
                "website": "lco",
                "tags": [
                        "web-dev",
                        "js"
                ]
        }
	`)

	var course course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("Valid")
		json.Unmarshal(jsonData, &course)
		fmt.Printf("%#v\n", course)
	} else {
		fmt.Println("Invalid")
	}

	// some case add data to key value pair without using struct

	var data map[string]interface{}
	json.Unmarshal(jsonData, &data)

	fmt.Printf("%#v\n", data)
}
