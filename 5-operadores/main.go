package main

import "fmt"

func main() {
	// ARITIMÉTICOS ==========
	soma := 1 + 2
	subtracao := 1 - 1
	divisao := 1 / 20
	multiplicacao := 1 * 2
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	var num1 int16 = 10
	var num2 int32 = 25

	//Você nao pode fazer alterações em  Tipos diferentes
	// soma := num1 + num2

	fmt.Println(num1, num2)
	// ===============

	//ATRIBUIÇÃO
	var var1 string = "String"
	var2 := "String 2"
	fmt.Println(var1, var2)
	//===============

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	//===============

	//OPERADORES LÓGICOS
	fmt.Println(false && true)
	fmt.Println(true || false)
	fmt.Println(!true || false)

	//OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero--
	numero += 15
	numero -= 15
	numero *= 3
	fmt.Println(numero)

	var texto string

	if numero > 5 {
		texto = "maior"
	} else {
		texto = "menor"
	}
	fmt.Println(texto)

}
