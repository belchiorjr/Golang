package main

import (
	"fmt"
)

func soma(numeros ...int) int{
	fmt.Println(numeros)

	total := 0 
	for _, numero := range numeros{
		total += numero
	}

	return total
}

func escrever (texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}


func main(){
	totalDaSoma  := soma(1,2,3,4,5,6,7,845,85)
	fmt.Println(totalDaSoma)


	escrever("Olá mundo ", 1,5,4,6,7,85,6,45)
}