package main

import (
	"fmt" 
)

func invertSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro (numero *int) {
	*numero = *numero * -1
}

func main() { 
	fmt.Println("Funções com retorno de Ponteiros")

	numero := 20 
	numerInvertido := invertSinal(numero)
	fmt.Println(numerInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}