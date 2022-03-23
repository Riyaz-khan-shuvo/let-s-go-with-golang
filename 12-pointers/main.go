package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of on pointers")

	// the default value of a pointer is nil

	var ptr *int
	fmt.Println("The value of Pointer is : ", ptr)

	myNumber := 25

	prt := &myNumber

	fmt.Println("The memory address is : ", prt)
	fmt.Println("The value of the address is : ", *prt)

	*prt = *prt * 3
	fmt.Println("New Value is : ", myNumber)

}

// pointer is a variable that stores the address of another variable.

// array is a special type of variable that stores multiple value in one single variable.
