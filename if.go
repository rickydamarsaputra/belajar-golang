package main

import "fmt"

func main() {
	name := "cahyo"

	if name == "ricky" && name != "damar" {
		fmt.Println("hello", name)
	} else {
		fmt.Println("who are you?")
	}
}
