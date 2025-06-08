package main

import (
	"errors"
	"fmt"
)

func main() {

	// integer types

	number := 1000000000000000000
	fmt.Println("The value of number is:", number)

	var number2 uint32 = 1000000000
	fmt.Println("The value of number2 is:", number2)

	var number3 rune = 1000000000
	fmt.Println("The value of number3 is:", number3)

	var number4 byte = 100
	fmt.Println("The value of number4 is:", number4)

	var number5 int
	fmt.Println("The value of number5 is:", number5)

	// real number types

	var realNumber1 float32 = 123.45
	fmt.Println("The value of realNumber is:", realNumber1)

	var realNumber2 float64 = 123.45
	fmt.Println("The value of realNumber2 is:", realNumber2)

	realNumber3 := 123.45
	fmt.Println("The value of realNumber3 is:", realNumber3)

	var realNumber4 float64
	fmt.Println("The value of realNumber4 is:", realNumber4)

	// string type
	var str string = "Hello, World!"
	fmt.Println("The value of str is:", str)

	str2 := "Hello, Go!"
	fmt.Println("The value of str2 is:", str2)

	character := 'A'
	fmt.Println("The value of character is:", character)

	var text string
	fmt.Println("The value of text is:", text)

	// boolean type
	var boolean1 bool
	fmt.Println("The value of boolean1 is:", boolean1)

	// error type
	var err error
	fmt.Println("The value of err is:", err)

	var err2 error = errors.New("This is an error")
	fmt.Println("The value of err2 is:", err2)
}
