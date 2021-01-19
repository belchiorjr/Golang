package main

import (
	"fmt" 
)


func main() { 
	fmt.Println("Programa Padrão de Concorrências Worker Pools")

	// Config canal de tarefas
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados) 
	go worker(tarefas, resultados) 
	go worker(tarefas, resultados) 
	go worker(tarefas, resultados) 

	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <- resultados
		fmt.Println(resultado)
	}
}
// tarefas <- canal que só recebe dados
// resultados -> canal que só envia os dados
func worker (tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci (posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}