package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the files in golang!!!")
	content := "I am learning Go lang."
	files, err := os.Create("./gotext.txt")
	checkNilError(err)
	length, err := io.WriteString(files, content)
	checkNilError(err)
	fmt.Println("Length is : ", length)
	defer files.Close()
	readFile("./gotext.txt")

}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("The data of this file is : ", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
