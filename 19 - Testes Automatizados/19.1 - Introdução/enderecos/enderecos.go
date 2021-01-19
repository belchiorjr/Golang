package enderecos

import (
	"strings"
)

// TipoDeEndereco - Função a ser exportada com iniciais maiuscula
func TipoDeEndereco (endereco string) string{
	tiposValidos := [] string {"rua", "avenida", "estrada", "rodovia"}
	enderecoEmLetrasMinusculas := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetrasMinusculas, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"
}