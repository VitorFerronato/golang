package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		int8
		int16
		int32
		int64
	*/

	number := 100000
	var number1 int32 = 10000

	fmt.Println(number)
	fmt.Println(number1)

	// Alias  INT32 = RUNE
	var number3 rune = 123456
	fmt.Println(number3)

	//Alias BYTE = UINT8
	var number4 byte = 123
	fmt.Println(number4)

	// FLOAT ======
	var float float32 = 123.45
	fmt.Println(float)

	var float1 float64 = 1234.556
	fmt.Println(float1)
	// ============

	// STRING ====
	var str string = "Texto"
	fmt.Println(str)

	str1 := "Texto 2"
	fmt.Println(str1)

	char := 'B'
	fmt.Println(char)
	// ============

	var texto string
	fmt.Println(texto)

	var boolean1 bool = true
	fmt.Println(boolean1)

	var erro error = errors.New("Erro interno")
	fmt.Println("ERRO", erro)

}
