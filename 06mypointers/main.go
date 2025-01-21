package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers!")

	myNumber := 20

	var ptr = &myNumber

	fmt.Println("Value of pointer", ptr)
	fmt.Println("Actual value of pointer", *ptr)

	*ptr += 2
	fmt.Println(*ptr)
}
