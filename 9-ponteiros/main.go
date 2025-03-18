package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	// Ponteiro é uma referência de memória
	var var3 int
	var ponteiro *int

	var3 = 100
	ponteiro = &var3

	fmt.Println(var3, ponteiro)

	// Arrays internos
	slice := make([]float32, 10, 15)
	fmt.Println(slice)
	fmt.Println(len(slice)) //length
	fmt.Println(cap(slice)) //capacidade

}
