package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Working with server!!!")
	PerformGetRequest()
}

func PerformGetRequest() {
	const url string = "http://localhost:8000/get"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var responseStrings strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseStrings.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseStrings.String())

	// fmt.Println(content)
	// fmt.Println(string(content))
}
