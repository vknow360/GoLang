//Declare three int variables num1, num2, and num3 and initialize them to any values. Calculate and print their average as a float64

package main

import "fmt"

func main() {
	var (
		num1 = 10
		num2 = 15
		num3 = 13
	)
	average := float64(num1+num2+num3) / 3

	fmt.Printf("The average is: %.2f\n", average)
}
