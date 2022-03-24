package main

import "fmt"

func main() {
	fmt.Println("Welcome the maps in golang!!!")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Rubey"
	languages["PY"] = "Python"
	fmt.Println(languages)
	delete(languages, "PY")
	fmt.Println(languages)

}
