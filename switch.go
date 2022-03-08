package main

import "fmt"

func main() {
	name := "candra"

	switch name {
	case "ricky":
		fmt.Println("hello", name)
	case "damar":
		fmt.Println("hello", name)
	default:
		fmt.Println("what is your name?")
	}

	switch number := 10; number >= 10 {
	case true:
		fmt.Println("lebih dari 10")
	case false:
		fmt.Println("kurang dari 10")
	}

	length := len(name)

	switch {
	case length > 5:
		fmt.Println("lebih dari 10")
	case length < 5:
		fmt.Println("kurang dari 10")
	}
}
