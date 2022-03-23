package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome learning slices !!!")
	var fruitsList = []string{"Apple", "Orange", "Banana"}

	fmt.Printf("The type of fruitsList is : %T\n", fruitsList)

	fmt.Println(fruitsList)

	// if the type of the variable is []string then we call it the slice.

	fruitsList = append(fruitsList, "tomato", "lichee")

	fmt.Println(fruitsList)
	fmt.Println(fruitsList[1:])
	fmt.Println(fruitsList[:3])
	fmt.Println(fruitsList[1:3])

	highScores := make([]int, 4)

	highScores[0] = 25
	highScores[1] = 55
	highScores[2] = 45
	highScores[3] = 35
	// highScores[4] = 35

	//  if we want to assing the value of highScores[4] = 35 it will give us error because the length of the slice is 4 but when we append the value it does not give the error because it is going to reallocate the memory

	highScores = append(highScores, 47, 29)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

}
