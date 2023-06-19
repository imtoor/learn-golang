package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Reynaldi",
		"address": "Makassar",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Reynaldy"
	book["oops"] = "salah"

	fmt.Println(len(book))
	fmt.Println(book)

	delete(book, "oops")

	fmt.Println(len(book))
	fmt.Println(book)

}
