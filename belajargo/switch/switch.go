package main

import "fmt"

func main() {
	name := "Dokja"

	switch name {
	case "Dokja":
		fmt.Println("HI Dokja")
		fmt.Println("HI Dokja")
	case "Hyuk":
		fmt.Println("HI Hyuk")
		fmt.Println("HI Hyuk")
	default:
		fmt.Println("HIII BLEH KENALAN?")
		fmt.Println("HIII BLEH KENALAN?")
	}
	dwitch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}
}
