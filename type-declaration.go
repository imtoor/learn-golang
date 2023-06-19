package main

import "fmt"

func main() {

	type NoKTP string
	type Married bool

	var noKtp NoKTP = "569182649812"
	var married Married = true

	fmt.Println(noKtp)
	fmt.Println(married)
}
