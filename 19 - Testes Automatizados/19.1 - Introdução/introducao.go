package main

import (
	"introducao-testes/enderecos"
	"fmt" 
)

func main() { 
	fmt.Println("Programa Introdução aos Testes Automatizados")

	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
}