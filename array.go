package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Reynaldi"
	names[1] = "Rahmat"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [3]int{
		90, 95, 80,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string
	fmt.Println(len(lagi))
}