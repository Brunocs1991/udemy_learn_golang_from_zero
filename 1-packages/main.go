package main

import (
	"fmt"
	"modulo/assistant"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Write from file main.go")
	assistant.Write()

	error:= checkmail.ValidateFormat("teste@gmail.com")
	fmt.Println("Email validation result:", error)
	fmt.Println("Email validation result:", checkmail.ValidateFormat("teste@"))
}
