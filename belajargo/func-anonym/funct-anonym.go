package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Your are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blacklist)
	registerUser("Najwa", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("Najwa", func(name string) bool {
		return name == "root"
	})
}
