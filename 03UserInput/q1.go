//Write a program that prompts the user to enter their first and last name. After the user enters their name, print a greeting message: "Hello, [First Name] [Last Name]"

package main

import "fmt"

func main() {
	var fName, lName string
	fmt.Scan(&fName, &lName)

	fmt.Println("Hello", fName, lName)

}
