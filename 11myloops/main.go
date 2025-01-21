package main

import "fmt"

func main() {
	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	// simple for loop
	for d := 0; d < len(days); d++ {
		fmt.Print(days[d] + " ")
	}
	fmt.Println("")

	//more advanced
	for d := range days {
		fmt.Print(days[d] + " ")
	}

	fmt.Println("")

	//for each
	for _, day := range days {
		fmt.Print(day + " ")
	}

	fmt.Println("")

	val := 1
	// while loop
	for val < 5 {
		fmt.Print(val, " ")
		val++
	}

}
