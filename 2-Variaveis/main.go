package main

import "fmt"

func main() {
	var var1 string = "Variavel 1"

	// Inferência de tipo
	var2 := "Variavel 2"

	fmt.Println("var1", var1)
	fmt.Println("var2", var2)

	var (
		var3 string = "var3"
		var4 string = "var4"
	)

	fmt.Println(var3, var4)

	var5, var6 := "Variavel 5", "Variável 6"

	fmt.Println(var5, var6)

	const const1 string = "Constante 1"

	fmt.Println(const1)

}
