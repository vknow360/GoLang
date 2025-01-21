package main

import "fmt"

func main() {
	sunny := Student{"Sunny", "vknow360@sunnythedeveloper.in", 19, true}
	fmt.Printf("%+v", sunny)
}

type Student struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
