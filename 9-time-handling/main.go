package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	// to formate the date and time

	fmt.Println(presentTime.Format("01-02-2006 , Monday"))

	// created date by ourself

	createdDate := time.Date(2000, time.February, 01, 10, 15, 17, 16, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:03:05 , Monday"))

}
