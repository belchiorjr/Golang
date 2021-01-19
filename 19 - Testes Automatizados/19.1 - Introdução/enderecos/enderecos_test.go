package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct{
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco (t *testing.T) {

	t.Parallel() // indica que esses testes podem rodar em paralelo

	cenariosDeTeste := [] cenarioDeTeste {
		{ "Rua ABC", "Rua"},
		{ "Avenida ABC", "Avenida"},
		{ "Rodovia dos Imigrantes", "Rodovia"},
		// { "Praça das Rosas", "Tipo Inválido"},
		{ "Estrada qualquer", "Estrada"},
		{ "RUA DAS VILAS", "Rua"},
		{ "AVENIDA REBOUÇAS", "Avenida"},
		// { "", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava '%s' e Recebeu '%s'", 
				cenario.retornoEsperado, 
				retornoRecebido,
			)
		}
	}
}

/////// go test ./... vai rodar todos testes do pacote
/////// go test -v  mostra as mensagens verbosas detalhadas
/// conbertura dos testes| go test --cover
/// gerar arquivo txt com o relatorio| go test --coverprofile resultado.txt
/// ler o arquivo resultado.txt| go tool cover --func=resultado.txt
/// ler o arquivo resultado.txt| go tool cover --html=resultado.txt