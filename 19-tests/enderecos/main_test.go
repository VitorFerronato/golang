package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestAdressType(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{
			"Rua avelino",
			"rua",
		},
		{
			"Avenida Paulista",
			"avenida",
		},
		{
			"Rua Santa cruz",
			"rua",
		},
		{
			"",
			"Invalid type",
		},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeCenarioRecebido := AdressType(cenario.enderecoInserido)
		if tipoDeCenarioRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoDeCenarioRecebido, cenario.retornoEsperado)
		}
	}
}
