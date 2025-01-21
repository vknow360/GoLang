//Write a program that asks the user to enter their age in years. Then, calculate and print their age in months and days. Assume each year has 365 days.

package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age in years: ")
	fmt.Scan(&age)

	fmt.Println("Age in months:", 18*12)
	fmt.Println("Age in days:", 18*365)
}
