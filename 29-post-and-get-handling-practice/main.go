package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to the practice part of get and post request!!!")
	// PerformGetRequest()
	PerformPostjsonRequest()
}

func PerformGetRequest() {

	const url = "http://localhost:8000/get"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var requestString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := requestString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(requestString.String())
}
func PerformPostjsonRequest() {
	const url = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	{
		"name":"Riyaz Hossain",
		"age":"22",
		"profession":"Student"
	}
	`)

	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
