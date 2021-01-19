
package main

import (
	"fmt"
)

func somar(numero1 int8, numero2 int8) int8 {
	return numero1+numero2
}


// função com multiplos retornos
func calculoMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}


func main(){
	soma := somar(10,20)
	fmt.Println(soma)

	// tipo funcao
	var f = func(text string) string {
		fmt.Println(text)
		return text +  " <<<< Return"
	}

	var resultSet = f("Parametro Text")
	fmt.Println(resultSet)

	somaNumeros, subtracaoNumeros := calculoMatematicos(10,50)
	fmt.Println(somaNumeros, subtracaoNumeros)

	somaNumerosOnly, _ := calculoMatematicos(10,50)
	fmt.Println(somaNumerosOnly)


}