package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	writeFile("myFile.txt", "Hello, world!")
	readFile("myFile.txt")
}

func writeFile(fileName string, content string) {
	file, err := os.Create("./" + fileName)
	if err != nil {
		panic(err)
	} else {
		length, err := io.WriteString(file, content)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fmt.Println("Length of file is", length)
	}
}

func readFile(fileName string) {
	bytes, err := os.ReadFile("./" + fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
