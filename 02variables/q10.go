//Declare a float64 variable height and initialize it with your height in meters (e.g., 1.75). Convert it to int and print both the original and converted values.

package main

import "fmt"

func main() {
	var height float64 = 1.75
	fmt.Println("Float value", height)
	fmt.Println("Integer value", int(height))
}
