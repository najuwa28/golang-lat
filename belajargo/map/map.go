package main

import "fmt"

func main() {

	person := map[string]string{
		"name":   "Najwa",
		"addres": "Jakarta",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["addres"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Jendoksi"
	book["author"] = "Sing n Song"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

}
