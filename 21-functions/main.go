package main

import "fmt"

func main() {
	fmt.Println("Welcome to the function !!!")
	greetings()
	greetingsTwo()

	var result = adder(5, 6)

	var manyAdder = proAdder(1, 2, 3, 4, 5)

	fmt.Println(manyAdder)

	fmt.Println(result)
}

func adder(valueOne int, valueTwo int) int {
	total := valueOne + valueTwo
	return total
}
func proAdder(values ...int) int {
	total := 0

	fmt.Println(values)

	for _, value := range values {
		total += value
	}
	return total
}

func greetingsTwo() {
	fmt.Println("This is Second Greetings!!!")
}

func greetings() {
	fmt.Println("This is First Greetings!!!")
}
