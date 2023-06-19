package main

import "fmt"

func main() {
	var val32 int32 = 32768
	var val64 int64 = int64(val32)
	var val16 int8 = int8(val32)

	fmt.Println(val32)
	fmt.Println(val64)
	fmt.Println(val16)

	var name = "Reynaldi Rahmat"
	var e = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
