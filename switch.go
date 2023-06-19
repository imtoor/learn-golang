package main

import "fmt"

func main() {

	name := "Reynaldy"

	switch name {
	case "Reynaldy":
		fmt.Println("Hello " + name)
	case "Rahmat":
		fmt.Println("Hello" + name)
	default:
		fmt.Println("Kamu bukan anggota")
	}

	switch length := len(name); length >= 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("panjang nama normal")
	}

	length := len(name)
	switch {
	case length >= 10:
		fmt.Println("nama terlalu panjang")
	case length >= 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("panjang nama normal")
	}

}
