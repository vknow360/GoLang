package main

import (
	"fmt"
	"strconv"
)

func main() {
	var PI float64 = 3.141597675757
	var i int = int(PI)
	fmt.Println("float value:", PI)
	fmt.Println("Integer value:", i)
	var f float64 = float64(i)
	fmt.Println("Again float value:", f)
	var s = strconv.FormatFloat(PI, 'f', -1, 64)
	println("String value:", s)
}
