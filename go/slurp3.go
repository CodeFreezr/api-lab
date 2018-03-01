package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type tag struct {
	Total int `json:"total"`
}

func main() {

	tagnames := [32]string{
		"graphviz", "d3.js", "svg",

		"html", "nodejs", "vue.js", "reactjs", "angular", "angularjs", "ember.js",

		"golang", "python", "c++", "c", "c#",

		"java", "groovy", "scala", "kotlin",

		"jenkins", "docker", "kubernetes", "git", "github", "gitlab",

		"sql", "nosql", "mongodb", "mysql", "postgresql", "oracle", "sqlite",
	}

	fmt.Println("No Tag-Name Total UnA NoA | UnA NoAPrc | NoSa NoSaPrc")

	for i, v := range tagnames {
		//fmt.Printf("%d = %v\n", i, v)
		s := v
		url1 := "https://api.stackexchange.com/2.2/questions?filter=total&tagged=" + s + "&site=stackoverflow"
		url2 := "https://api.stackexchange.com/2.2/questions/unanswered?filter=total&tagged=" + s + "&site=stackoverflow"
		url3 := "https://api.stackexchange.com/2.2/questions/no-answers?filter=total&tagged=" + s + "&site=stackoverflow"

		ttl := float64(getNum(url1)) //total
		una := float64(getNum(url2)) //unanswered
		noa := float64(getNum(url3)) //noanswers

		unaproc := (una / ttl) * 100 //unanswered% vs. totals
		noaproc := (noa / ttl) * 100 //noanswers% vs. totals

		nosa := una - noa // nosatisfaction
		nosaproc := (nosa / una) * 100

		fmt.Println(i+1, s, ttl, una, noa, "|", unaproc, noaproc, "|", nosa, nosaproc)

		time.Sleep(1000)
	}
}

func add(x int, y int) int {
	return x + y
}

func getNum(url string) int {

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	tag1 := tag{}
	jsonErr := json.Unmarshal(body, &tag1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return tag1.Total
}
