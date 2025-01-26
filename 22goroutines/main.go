package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// go greet("Hello")
	// greet("World")

	websiteList := []string{
		"https://google.com",
		"https://go.dev",
		"https://facebook.com",
		"https://instagram.com",
	}

	for _, url := range websiteList {
		go getStatusCode(url)
		wg.Add(1)
	}
	wg.Wait()
}

func greet(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(str)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	response, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%d status code for %s\n", response.StatusCode, endpoint)
	}
}

func createFile() {
	file, _ := os.Create("myFile.txt")
	file.WriteString("Hello, World!")
	file.Close()
}
