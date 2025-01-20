// Multiple Variable Declaration

package main

import "fmt"

func main() {
	var (
		a = 10
		b = 20
	)
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", b, a, b-a)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
