package main

import "fmt"

func main() {
	fmt.Println("Defear Practice ")
	myDefear()

	defer fmt.Println("World!")
	fmt.Println("Hello")
}

func myDefear() {
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
