package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome  remove a value from slices !!!")

	// how to remove a value from slices base on index

	var index int = 2

	var courses = []string{"JavaScript", "ReactJs", "swift", "Python", "Django"}
	fmt.Println(courses)

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
