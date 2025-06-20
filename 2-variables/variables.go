package main

import "fmt"

func main() {
	var var1 string = "var1"
	fmt.Println(var1)

	var2 := "var2"
	fmt.Println(var2)

	var (
		var3 = "var3"
		var4 = "var4"
	)

	fmt.Println(var3, var4)

	var5, var6 := "var5", "var6"
	fmt.Println(var5, var6)

	const const1 string = "const1"
	fmt.Println(const1)

	var5, var6 = var6, var5 // Swapping values
	fmt.Println(var5, var6)

}
