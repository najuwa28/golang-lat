package main

import "fmt"

func main() {

	var result = 15 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	var i = 10
	i += 15 // i = i + 10 augmented

	fmt.Println(i)

	i++ // i = i +1
	fmt.Println(i)

	var negative = -100
	var positive = 100
	fmt.Println(negative)
	fmt.Println(positive)

	var value1 = 10
	value1 += 100
	fmt.Println(value1) //karna value1 = isi + isi previous
}
