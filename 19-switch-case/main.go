package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in golang")
	rand.Seed(time.Now().UnixNano())
	discNumber := rand.Intn(6) + 1
	fmt.Println(discNumber)

	switch discNumber {
	case 1:
		fmt.Println("Dice value is one and you can open!!!")
	case 2:
		fmt.Println("You have move 2 spot")
	case 3:
		fmt.Println("You have move to 3 spot")
		fallthrough
	case 4:
		fmt.Println("You have move to 4 spot")
		fallthrough
	case 5:
		fmt.Println("You have move to 5 spot")
	case 6:
		fmt.Println("You have move to 6 spot")

	}

}
