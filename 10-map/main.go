package main

import "fmt"

func main() {
	// Chaves tem que ser do mesmo tipo e as chaves também
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro":  "vitor",
			"sobrenome": "Teste",
		},
	}

	fmt.Println(usuario, usuario2)
	delete(usuario, "nome")
	fmt.Println(usuario)

	usuario2["signo"] = map[string]string{
		"nome": "Virgem",
	}

}
