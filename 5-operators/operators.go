package main

func main() {
	// Operators in Go
	// Arithmetic Operators: +, -, *, /, %
	// Comparison Operators: ==, !=, <, >, <=, >=
	// Logical Operators: &&, ||, !
	// Bitwise Operators: &, |, ^, <<, >>
	// Assignment Operators: =, +=, -=, *=, /=, %=

	a := 10
	b := 5

	sum := a + b
	diff := a - b
	prod := a * b
	quot := a / b
	mod := a % b

	println("Sum:", sum)
	println("Difference:", diff)
	println("Product:", prod)
	println("Quotient:", quot)
	println("Modulus:", mod)

	// Comparison
	isEqual := a == b
	isNotEqual := a != b
	isGreater := a > b
	isLess := a < b
	isGreaterOrEqual := a >= b
	isLessOrEqual := a <= b
	println("Is Equal:", isEqual)
	println("Is Not Equal:", isNotEqual)
	println("Is Greater:", isGreater)
	println("Is Less:", isLess)
	println("Is Greater or Equal:", isGreaterOrEqual)
	println("Is Less or Equal:", isLessOrEqual)

	// Logical
	isTrue := true
	isFalse := false
	logicalAnd := isTrue && isFalse
	logicalOr := isTrue || isFalse
	logicalNot := !isTrue
	println("Logical AND:", logicalAnd)
	println("Logical OR:", logicalOr)
	println("Logical NOT:", logicalNot)
	// Bitwise
	bitwiseAnd := a & b
	bitwiseOr := a | b
	bitwiseXor := a ^ b
	bitwiseLeftShift := a << 1

	bitwiseRightShift := a >> 1
	println("Bitwise AND:", bitwiseAnd)
	println("Bitwise OR:", bitwiseOr)
	println("Bitwise XOR:", bitwiseXor)
	println("Bitwise Left Shift:", bitwiseLeftShift)
	println("Bitwise Right Shift:", bitwiseRightShift)

	// Assignment
	a += b
	println("After Assignment, a:", a)
	b -= 2
	println("After Assignment, b:", b)
	println("Final Values - a:", a, "b:", b)
	println("End of Operators Example")
}