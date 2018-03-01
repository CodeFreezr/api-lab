//go get -u gopkg.in/resty.v1
package main

import (
	"fmt"

	"gopkg.in/resty.v1"
)

type tag2 struct {
	Total int `json:"total"`
}

var url = "https://api.stackexchange.com/2.2/questions?filter=total&tagged=graphviz&site=stackoverflow"

//var url = "http://httpbin.org/get"

func main() {
	fmt.Println("Hello, 世界")
	resp, err := resty.R().Get(url)

	var body = resp.String

	tag1 := tag{}
	//jsonErr := json.Unmarshal(resp, &tag1)
	fmt.Printf("\nError: %v", err)

	fmt.Printf("\nResponse Body: %v", resp)
}
