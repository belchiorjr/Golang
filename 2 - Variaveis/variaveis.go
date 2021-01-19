package main

import (
	"fmt"
)

func main() {
	var varialvel1 string = "Variável 1"
	variavel2 := "variável 2"
	
	fmt.Println(varialvel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)


	variavel5, variavel6 := "Variavel 5", "Variavel 6"


	fmt.Println(variavel5)
	fmt.Println(variavel6)

	const constante1 string = "Constante 1"

	fmt.Println(constante1)

	// Troca de valores entre variáveis
	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5 + " (Variavel 5 trocada com a 6)")
	fmt.Println(variavel6 + " (Variavel 6 trocada com a 5)")

}