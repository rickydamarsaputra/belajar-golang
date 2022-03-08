package main

import "fmt"

func main() {
	var numbers = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var numberSlice = numbers[1:4]
	var numberSlice2 = append(numberSlice, 30)

	numbers[1] = 50

	fmt.Println(numberSlice)
	fmt.Println(len(numberSlice))
	fmt.Println(cap(numberSlice))

	fmt.Println(numberSlice2)

	newSlice := make([]int, 5)

	newSlice[0] = 10
	newSlice[1] = 20
	newSlice[2] = 30
	newSlice[3] = 40
	newSlice[4] = 50

	fmt.Println(newSlice)

	copySlice := make([]int, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	copySlice = append(copySlice, 60)

	fmt.Println(copySlice)

	array := [3]int{1, 2, 3}
	slice := []int{1, 2, 3}

	fmt.Println(array)
	fmt.Println(slice)
}
