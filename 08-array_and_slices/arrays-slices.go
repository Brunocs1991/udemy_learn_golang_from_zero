package main
import "fmt"
func main() {
	// Declare and initialize an array of integers
	numbers := [5]int{1, 2, 3, 4, 5}

	// Print the entire array
	fmt.Println("Array:", numbers)

	// Access elements in the array
	fmt.Println("First element:", numbers[0])
	fmt.Println("Second element:", numbers[1])

	// Modify an element in the array
	numbers[2] = 10
	fmt.Println("Modified Array:", numbers)

	// Declare and initialize a slice of strings
	fruits := []string{"apple", "banana", "cherry"}

	// Print the entire slice
	fmt.Println("Slice:", fruits)

	// Access elements in the slice
	fmt.Println("First fruit:", fruits[0])
	fmt.Println("Second fruit:", fruits[1])

	// Append an element to the slice
	fruits = append(fruits, "date")
	fmt.Println("Slice after append:", fruits)
}