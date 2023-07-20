package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}
func (a Customer) sayHuuu() {
	fmt.Println("Huuuu from", a.Name)
}
func main() {
	var najwa Customer
	najwa.Name = "Najwa"
	najwa.Address = "Gorda"
	najwa.Age = 17

	najwa.sayHi("Aurel")
	najwa.sayHuuu()

	/** fmt.Println(najwa.Name)
	fmt.Println(najwa.Address)
	fmt.Println(najwa.Age)

	aurel := Customer{
		Name:    "Aurel",
		Address: "Fuji",
		Age:     17,
	}
	fmt.Println(aurel)

	lia := Customer{"Naju", "Makmur", 17}
	fmt.Println(lia) */
}
