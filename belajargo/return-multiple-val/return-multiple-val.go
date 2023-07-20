package main

import "fmt"

func getFullName() (string, string, string) {
	return "Najwa", "Aurelia", "Balqist"
}

func main() {
	firstName, middName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(middName)
	fmt.Println(lastName)
}
