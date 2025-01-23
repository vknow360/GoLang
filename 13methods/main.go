package main

import "fmt"

func main() {
	sunny := Student{"Sunny", "vknow360@sunnythedeveloper.in", 19, true}
	fmt.Printf("%+v\n", sunny)
	sunny.GetStatus()
}

type Student struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (s Student) GetStatus() {
	fmt.Println("Is active:", s.Status)
}
