package main

import (
	"fmt" 
)

func main() { 
	fmt.Println("Funções Anônimas!")

	func () {
			fmt.Println("Função Anônima Informada!")
	} () // Os parenteses inicializa a função anônima

	func (texto string) {
			fmt.Println(texto)
	} ("Teste de impressão") // Inicializa pasando paremtros

	retorno := func (texto string) string {
		return 	fmt.Sprintf("Recebido -> %s", texto)
	} ("Teste com Sprintf")

	fmt.Println(retorno)
	
}