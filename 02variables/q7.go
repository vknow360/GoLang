// Declare two bool variables x and y. Initialize them with different boolean values. Print the result of x && y (logical AND) and x || y (logical OR)
package main

import "fmt"

func main() {
	var (
		x = false
		y = true
	)
	fmt.Printf("x && y: %t\n", x && y)
	fmt.Printf("x || y: %t", x || y)
}
