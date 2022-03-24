package main

import "fmt"

func main() {
	fmt.Println("Welcome to the defear in golang")
	myDefear()
	defer fmt.Println("world")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

}

func myDefear() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}
