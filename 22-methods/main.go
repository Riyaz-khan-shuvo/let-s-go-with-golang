package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Methods in golang!!!")

	riyaz := User{"Riyaz Hossain", "mdriyaz5965@gmail.com", 17, true}
	fmt.Println(riyaz)
	fmt.Printf("Riyaz Details are : %+v\n", riyaz)
	riyaz.getStatus()
	riyaz.setNewMail()
	fmt.Printf("Riyaz Details are : %+v\n", riyaz)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) getStatus() {
	fmt.Println("Is User Active : ", u.Status)
}

func (u User) setNewMail() {
	u.Email = "riyazhossain5965@gmail.com"
	fmt.Println("New Mail is : ", u.Email)
}
