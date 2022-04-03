package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome , Working with file !!!")
	content := "Riyaz Khan Shuvo"
	files, err := os.Create("./hello.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(files, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is : ", length)
	defer files.Close()
	readFile("./hello.txt")
}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("The file data is : ", databyte)
	fmt.Println("The file data is : ", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
