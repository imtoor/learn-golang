package main

import "fmt"

func main() {

	// counter := 0

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke: ", counter)
	// 	counter++
	// }

	// init dan postStatement

	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke: ", counter)
	// }

	// slice := []string{"Reynaldy", "Rahmat", "Ganteng"}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Print(slice[i] + " ")
	// }

	names := []string{"Reynaldy", "Rahmat", "Ganteng"}

	// for index, name := range names {
	// 	fmt.Println("Index ", index, " = ", name)
	// }

	for _, name := range names {
		fmt.Println(name)
	}

	person := make(map[string]string)
	person["name"] = "Reynaldy Rahmat"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, " = ", value)
	}

}
