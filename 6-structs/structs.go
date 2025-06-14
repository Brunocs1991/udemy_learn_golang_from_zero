package main

import "fmt"

// Define a struct for a person
	type Person struct {
		name string
		age  int
		address Address // Embedded struct
	}
	// Define a struct for a car
	type Car struct {
		make  string
		model string
		year  int
	}

	type Address struct {
		street string
		city   string
		state  string
		zip   string
	}

func main() {


	// Create a new person
	person := Person{
		name: "Alice",
		age:  30,
		address: Address{
			street: "123 Main St",
			city:   "Springfield",
			state:  "IL",
			zip:    "62701",
		},
	}

	// Print the person's details
	fmt.Println(person)
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)

	// Create a new car
	car := Car{
		make:  "Toyota",
		model: "Corolla",
		year:  2020,
	}

	// Print the car's details
	fmt.Println(car)
	fmt.Println("Car Make:", car.make)
	fmt.Println("Car Model:", car.model)
	fmt.Println("Car Year:", car.year)
}