package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Najwa", "Aurelia", "Balqist", "Dokja", "Hyuk"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(Slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Najwa"
	person["title"] = "Progay"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
