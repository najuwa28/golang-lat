package main

import "fmt"

func main() {

	type NoKTP string
	type Married bool

	var ktpNajwa NoKTP = "22242772"
	var marriageStatus Married = false

	fmt.Println(ktpNajwa)
	fmt.Println(marriageStatus)

}
