package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const murl = "http://localhost:8000/post"

func main() {
	//PerformGet(murl)
	//PerformJsonPostRequest()
	PerformFormDataPostRequest()
}

func PerformGet(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformJsonPostRequest() {
	requestBody := strings.NewReader(`
		{
			"courseName" : "Let's go with GoLang",
			"price": 0,
			"platform" : "Youtube"
		}
	`)

	response, err := http.Post(murl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformFormDataPostRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "Sunny")
	data.Add("lastname", "Gupta")
	data.Add("email", "vknow360@gmail.com")

	response, err := http.PostForm(murl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
