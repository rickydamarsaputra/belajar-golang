package main

import "fmt"

func main() {
	var numbers [3]int

	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3

	fmt.Println(numbers)

	var numbers2 = [3]int{1, 2, 3}

	fmt.Println(numbers2)

	fmt.Println(len(numbers))
	fmt.Println(numbers[0])
}
