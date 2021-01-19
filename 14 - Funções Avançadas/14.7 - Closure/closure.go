package main

import (
	"fmt" 
)


func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() { 
	fmt.Println("Funções Closure")

	funcaoRecebida := closure()
	funcaoRecebida()
}