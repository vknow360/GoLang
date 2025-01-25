package main

import (
	"fmt"
	"github/vknow360/GoLang/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("Server is getting started...")
	r := router.Router()
	http.ListenAndServe(":4000", r)
	fmt.Println("Listening at port 4000...")
}
