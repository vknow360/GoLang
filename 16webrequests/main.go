package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://example.com"

func main() {
	fmt.Println(get(url))
}

func get(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code:", response.StatusCode)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}
