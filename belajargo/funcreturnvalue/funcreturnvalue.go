package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello brou"
	} else {
		return "Hello " + name
	}
}
func main() {
	result := getHello("Najwa")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
