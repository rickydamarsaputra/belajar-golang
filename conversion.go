package main

import "fmt"

func main() {
	var number8 int8 = 100
	var number32 int32 = int32(number8)
	var name string = "ricky damar"
	var name1 string = string(name[0])

	fmt.Println(number8)
	fmt.Println(number32)
	fmt.Println(name)
	fmt.Println(name1)
}
