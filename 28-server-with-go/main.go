package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Hello , we are creating the server ")
	PerformGetRequest()

}

func PerformGetRequest() {
	const url string = "http://localhost:8000/get"
	response, err := http.Get(url)
	getError(err)
	defer response.Body.Close()

	fmt.Println("Status Code is : ", response.Status)
	fmt.Println("Content Length is : ", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	ByteCount, _ := responseString.Write(content)
	fmt.Println(ByteCount)
	fmt.Println(responseString.String())

	// or

	fmt.Println("The strings is : ", string(content))

}

func getError(err error) {
	if err != nil {
		panic(err)
	}
}
