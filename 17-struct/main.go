package main

import "fmt"

func main() {
	fmt.Println("About Struct")

	// in go there is no inheritance ; no super or parent

	riyaz := User{"Riyaz", "mdriyaz@go.dev", true, 19}
	fmt.Println(riyaz)
	fmt.Printf("The Details of Riyaz : %+v", riyaz)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
