package main

import "fmt"

type usuario struct {
	name     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.name = "Davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua não sei", 0}

	u2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(u2)

	u3 := usuario{idade: 22}
	fmt.Println(u3)

}
