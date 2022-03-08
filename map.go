package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "ricky",
		"address": "surabaya",
	}

	person["hobby"] = "watch anime"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(len(person))

	delete(person, "hobby")

	fmt.Println(person)
	fmt.Println(len(person))
}
