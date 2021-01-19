package main

import (
	"time"
	"fmt" 
)

func main() { 
	// CONCORRÊNCIA != PARALELISMO
	fmt.Println("Programa Canais")

	canal := make(chan string)

	go escrever("Olá Mundo", canal)
	
	fmt.Println("Depois da função escrever começar a ser executada!")

	// Opcao com o flag aberto
	// for{
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa!")

}


func escrever (texto string, canal chan string) {
	//time.Sleep(time.Second * 5)
	for i:=0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	// Fechar o canal
	close(canal)
}