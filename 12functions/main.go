package main

import "fmt"

func main() {
	// result := add(3, 5)
	// fmt.Println("Result is", result)

	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of numbers is", sumOfNumbers(s...))

}

// func add(num1 int, num2 int) int {
// 	return num1 + num2
// }

func sumOfNumbers(slice ...int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	return sum
}
