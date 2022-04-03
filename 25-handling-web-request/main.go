package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://riyaz-khan-shuvo.netlify.app/"

func main() {
	fmt.Println("Handle web request ")

	response, err := http.Get(url)
	checkNilError(err)
	fmt.Printf("Respose Type is : %T \n", response)
	defer response.Body.Close()
	databye, err := ioutil.ReadAll(response.Body)
	checkNilError(err)
	fmt.Println(string(databye))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
