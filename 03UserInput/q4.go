//Write a program that asks the user to enter a number and checks whether the number is even or odd. Print "Even" or "Odd" based on the input.

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

}
