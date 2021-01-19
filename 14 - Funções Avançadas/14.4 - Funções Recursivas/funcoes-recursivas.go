package main

import (
	"fmt" 
)

func fibonacci (posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() { 
	fmt.Println("Funções Recursivas")

	// Fibonacci
	
	posicao := uint(12)
	fmt.Println(fibonacci(posicao))
	
	posicao2 := uint(15)
	fmt.Println(fibonacci(posicao2))
	
	posicao3 := uint(1)
	fmt.Println(fibonacci(posicao3))

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}