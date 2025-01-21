package main

import "fmt"

func main() {
	var fruitlist [4]string
	fruitlist[0] = "Apple"
	fruitlist[1] = "Banana"
	fruitlist[2] = "Guvava"
	fruitlist[3] = "Orange"

	fmt.Println(fruitlist)
	fmt.Println(len(fruitlist))

	fmt.Printf("%T\n", fruitlist)

	var vegList = [5]string{"potato", "cabbage", "raddish"}
	fmt.Println(vegList)
}
