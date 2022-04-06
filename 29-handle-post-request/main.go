package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome , Handle Post Request!!!")

	// PerformGetRequest()

	PerformPostjsonRequest()
}

func PerformGetRequest() {
	const url string = "http://localhost:8000/get"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}

func PerformPostjsonRequest() {

	const url = "http://localhost:8000/post"

	// fake json playload

	requestBody := strings.NewReader(`
		{
			"courseName":"Let's go with golang",
			"price":"0",
			"platform":"riyaz-khan-shuov.netlify.app"
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
