package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := range 10 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"Vitor", "Eduardo", "Ferronato"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
