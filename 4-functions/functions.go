package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func calculate(a, b int) (int, int) {
	add := a + b
	subtract := a - b
	return add, subtract
}

func main() {
	resultAdd := add(10, 5)
	fmt.Println("Addition Result:", resultAdd)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("Hello, World!")
	fmt.Println("Function Result:", result)

	resultAdd, resultSubtract := calculate(20, 10)
	fmt.Println("Addition Result:", resultAdd)
	fmt.Println("Subtraction Result:", resultSubtract)

	resultAdd2, _ := calculate(30, 15)
	fmt.Println("Addition Result 2:", resultAdd2)

	_, resultSubtract2 := calculate(50, 25)
	fmt.Println("Subtraction Result 2:", resultSubtract2)
}
