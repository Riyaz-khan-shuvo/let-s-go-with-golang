package main

import "fmt"

func main() {
	fmt.Println("Welcome to the loops in go !!!")
	days := []string{"Sunday", "Monday", "WednesDay", "Friday"}

	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// same problem using for range loop

	for i, day := range days {
		fmt.Println("The index is : ", i, "The day is : ", day)
	}

}
