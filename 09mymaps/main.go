package main

import (
	"fmt"
)

func main() {

	languages := make(map[int]string)
	languages[0] = "Java"
	languages[1] = "C"
	languages[2] = "C++"

	fmt.Println(languages[2])

	for key, value := range languages {
		fmt.Printf("%v : %v\n", key, value)
	}
}
