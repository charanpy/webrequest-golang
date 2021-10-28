package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

const myurl = "https://jsonplaceholder.typicode.com/posts"

type post struct {
	UserId      int    `json:"userId"`
	Id          int    `json:"id"`
	Title       string `json:"title"`
	PostContent string `json:"post"`
}

func main() {
	// GetRequest(myurl)
	// PostRequest()
	buildQueryParams()
}

func loopSlice(slice []post) {
	for key, value := range slice {
		fmt.Println(key, value.Title)
	}
}

func GetRequest(url string) {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	// fmt.Println(string(content))
	if err != nil {
		panic(err)
	}

	isValid := json.Valid(content)

	var post []post

	if isValid {
		json.Unmarshal(content, &post)
		fmt.Printf("%#v\n", post)
	}

	loopSlice(post)
	fmt.Println(url)
}

func PostRequest() {
	rand.Seed(time.Now().UnixNano())

	var post post

	var UserId = rand.Intn(20)
	var Id = rand.Intn(100)

	post.UserId = UserId
	post.Id = Id

	var Title string
	fmt.Println("Enter title")
	fmt.Scanf("%s", &Title)
	post.Title = Title

	var PostContent string
	fmt.Println("Enter content")
	fmt.Scanf("%s", &PostContent)
	post.PostContent = PostContent

	data, err := json.Marshal(post)

	if err != nil {
		panic(err)
	}

	fmt.Println(data)

	response, err := http.Post(myurl, "application/json", bytes.NewBuffer(data))

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("success", string(content))
}

func buildQueryParams() {
	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "jsonplaceholder.typicode.com",
		Path:     "posts",
		RawQuery: "postId=1",
	}

	stringifiedURL := partsOfURL.String()
	fmt.Println(stringifiedURL)
	GetRequest(stringifiedURL)
}
