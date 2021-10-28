package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web req")
	// PerformGetRequest()
	// PerformPostJSONRequest()
	// PerformPostFormDataRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code", response.StatusCode)
	fmt.Println("Content len", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	// another way to print res body
	// var responseString strings.Builder
	// byteCount, _ := responseString.Write(content)

	// fmt.Println(byteCount)
	// fmt.Println(responseString.String())

	fmt.Println(string(content))
}

func PerformPostJSONRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"course":"react",
			"price":500
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormDataRequest() {
	const myurl = "http://localhost:8000/postform"

	// formdata

	data := url.Values{}
	data.Add("name", "charan")
	data.Add("favLanf", "js")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
