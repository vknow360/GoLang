//Write a program that takes two integers as input from the user and calculates and prints their sum.

package main

import "fmt"

func main() {
	var (
		a int
		b int
	)
	fmt.Scan(&a, &b)
	fmt.Println("a+b =", a+b)
}
