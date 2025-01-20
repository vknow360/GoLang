//Declare an int variable num and initialize it to any number. Print whether the number is even or odd using the modulo operator (%).

package main

import "fmt"

func main() {
	var num = 776

	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
