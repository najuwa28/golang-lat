package main

import "fmt"

func main() {
	fmt.Println("Najwa")
	fmt.Println("Najwa Aurelia")
	fmt.Println("Najwa Aurelia Balqist")

	fmt.Println(len("Najwa"))

	var name = "Najwa Aurelia"
	var e = name[0]
	var eString = string(e) //konversi string byte e

	fmt.Println(name)
	fmt.Println(eString)

}
