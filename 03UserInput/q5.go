/*Write a program that asks the user to enter a temperature in Celsius*/

package main

import "fmt"

func main() {
	var temp int // Celsius
	fmt.Scan(&temp)

	fmt.Println("Temperature in Fahrenheit: ", temp*9/5+32)

}
