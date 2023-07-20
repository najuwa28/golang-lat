package main

import "fmt"

func main() {
	name := "Najwa"
	counter := 0

	increment := func() {
		name := "Aurelia"
		fmt.Println("Increment")
		fmt.Println(name)
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
