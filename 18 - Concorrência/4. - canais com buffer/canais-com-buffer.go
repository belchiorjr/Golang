package main

import (
	"fmt" 
)

func main() { 
	fmt.Println("Programa Canais com Buffer")

	canal := make(chan string, 2) // especificar com Buffer
	canal <- "Olá Mundo"
	canal <- "Programando em GO!"
	//canal <- "Aqui vai acontercer o DeadLock" // Acontece o Dead lock porque estoura o buffer
	
	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)

	//close(canal)

}