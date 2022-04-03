package main

import (
	"fmt"
	"io"
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

}
