//Declare two string variables, firstName and lastName. Concatenate them to form a full name and print it.

package main

import "fmt"

func main() {
	var firstName, lastName = "Sunny", "Gupta"
	fullName := firstName + " " + lastName
	fmt.Println(fullName)
}
