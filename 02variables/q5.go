package main

import (
	"fmt"
	"strconv"
)

func main() {
	var age int = 18
	fmt.Println(age)
	fmt.Printf("type of age: %T \n", age)
	var sAge string = strconv.Itoa(age)
	fmt.Println(sAge)
	fmt.Printf("type of age: %T \n", sAge)
}
