package main

import "fmt"

func main() {
	// In Go, there is no direct support for inheritance like in some other languages.
	// However, we can achieve similar behavior using struct embedding.

	// Define a base struct
	type Animal struct {
		Name string
	}

	// Define a derived struct that embeds the base struct
	type Dog struct {
		Animal // Embedding the Animal struct
		Breed  string
	}

	// define a derived struct that embeds another struct
	type Cat struct {
		Animal // Embedding the Animal struct
		Color  string
	}

	// Create an instance of Dog
	myDog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	// Access fields from the embedded struct
	fmt.Println(myDog)
	fmt.Println("Dog Name:", myDog.Name)
	fmt.Println("Dog Breed:", myDog.Breed)

	// Create an instance of Cat
	myCat := Cat{
		Animal: Animal{Name: "Whiskers"},
		Color:  "Tabby",
	}
	// Access fields from the embedded struct
	fmt.Println(myCat)
	fmt.Println("Cat Name:", myCat.Name)
	fmt.Println("Cat Color:", myCat.Color)
	// Note: This is not true inheritance, but rather a way to achieve similar functionality
	// by embedding structs. Go promotes composition over inheritance.
	// This allows us to reuse code and create more complex types without the need for traditional inheritance.
}