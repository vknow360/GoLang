package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fruitList = append(fruitList, "Banana", "Guvava", "Mango")
	fmt.Println(fruitList)

	fruitList = fruitList[1:5]
	fmt.Println(fruitList)

	highScores := make([]int, 2)
	highScores[0] = 5
	highScores[1] = 3

	highScores = append(highScores, 1, 2, 4, 6, 7, 8)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

}
