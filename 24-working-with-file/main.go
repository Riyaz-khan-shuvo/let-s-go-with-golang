package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to the files in golang!!!")

	content := "I am learning Go lang."
	files, err := os.Create("./gotext.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(files, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is : ", length)

	defer files.Close()

}
