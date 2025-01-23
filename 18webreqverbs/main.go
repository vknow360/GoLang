package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://localhost:8000"

func main() {
	PerformGet(url)
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
