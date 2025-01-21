package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study!")

	presentDate := time.Now()
	fmt.Println(presentDate)

	fmt.Println(presentDate.Format("02-01-2006 Monday"))

	createdDate := time.Date(2006, time.May, 10, 4, 15, 0, 0, time.Now().UTC().Location())
	fmt.Println(createdDate)
}
