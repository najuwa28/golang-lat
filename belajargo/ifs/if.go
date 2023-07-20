package main

import "fmt"

func main() {

	var name = "Dokja"

	if name == "Dokja" {
		fmt.Println("Hi Dokja")
	} else if name == "Hyuk" {
		fmt.Println("HIII Hyuk")
	} else if name == "Najwa" {
		fmt.Println("HIII NAJWA")
	} else {
		fmt.Println("HIII BLEH KENALAN?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
		fmt.Println(length)
	}

}
