package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://sunnythedeveloper.in/extension-buying-policy/"

func main() {
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "sunnythedeveloper.in",
		Path:     "contact",
		RawQuery: "user=admin",
	}

	fmt.Println(partsOfUrl.String())
}
