package main

import "fmt"

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. Resutlado será retornado!")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false

}

func main() {
	alunoEstaAprovado(5.5, 4)
}
