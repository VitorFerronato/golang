package enderecos

import (
	"strings"
)

// AdressType Verifica se o endereço tem um tipo válido
func AdressType(endereco string) string {
	validAdress := []string{"rua", "avenida", "estrada", "rodovia"}

	adressToLowerCase := strings.ToLower(endereco)

	firstWord := strings.Split(adressToLowerCase, " ")[0]

	hasAValidType := false

	for _, tipo := range validAdress {

		if tipo == firstWord {
			hasAValidType = true
		}
	}

	if hasAValidType {
		return strings.Title(firstWord)
	}

	return "Invalid type"
}
