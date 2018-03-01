//go get -u gopkg.in/resty.v1
package main

import (
	"fmt"

	"gopkg.in/resty.v1"
)

type total struct {
	Total int `json:"total"`
}

var url = "https://api.stackexchange.com/2.2/questions?filter=total&tagged=graphviz&site=stackoverflow"

//var url = "http://httpbin.org/get"

func main() {
	fmt.Println("Hello, 世界")
	resp, err := resty.R().Get(url)
	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Body: %v", resp)
}
