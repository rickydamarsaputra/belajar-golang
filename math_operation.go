package main

import "fmt"

func main() {
	num1 := 10
	num2 := 5

	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)

	num1 += 5
	fmt.Println(num1)
	num1 -= 5
	fmt.Println(num1)
	num1 *= 5
	fmt.Println(num1)
	num1 /= 5
	fmt.Println(num1)
	num1 %= 5
	fmt.Println(num1)
}
