package main

import (
	"fmt"
	"tests/enderecos"
)

func main() {
	adressType := enderecos.AdressType("Rua santa cruz")
	fmt.Println(adressType)
}
