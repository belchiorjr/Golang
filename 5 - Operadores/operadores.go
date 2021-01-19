package main

import (
	"fmt"
)


func main() {
	// Aritméticos + - * / %

		soma :=  1+2
		subtracao := 1-2
		divisao := 10/4
		multiplicacao := 10 * 50
		restoDaDivisao := 10%2

		fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)


	// Atribuição = := 
	var variavelExplicitandoTipo string = "Variável com tipo explicitado no código"
	var variavelImplicitandoTipo string = "Variável com tipo implicitando no código"
	
	fmt.Println(variavelExplicitandoTipo, variavelImplicitandoTipo)


	// Operações Relacionais
	fmt.Println(1>2, 2<1, 2==1, 2!=1)


	// Operadores lógicos
	fmt.Println(true && false, true || false, true == false)

	// Operadores unários
	numero := 10
	numero++
	fmt.Println(numero)

	
	numero+=50
	fmt.Println(numero)


	/// Ternário não tem


}