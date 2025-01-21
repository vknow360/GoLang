package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	var tmp, _ = reader.ReadString('\n')
	age, _ := strconv.Atoi(strings.TrimSpace(tmp))
	fmt.Printf("%s is %d years old", name, age)

}
