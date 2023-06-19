package main

import "fmt"

func main() {

	var nama = "Aldi"

	if nama == "Reynaldy" {
		fmt.Println("Hello " + nama)
	} else if nama == "Rahmat" {
		fmt.Println("Hello " + nama)
	} else {
		fmt.Println("Hello " + nama + ", kamu bukan anggota grup!")
	}

	if length := len(nama); length >= 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Panjang nama normal")
	}

}
