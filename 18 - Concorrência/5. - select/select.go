package main

import (
	"time"
	"fmt" 
)

func main() { 
	fmt.Println("Programa Select")

	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for{
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for{
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()


	for {
		// Assim  perca de tempo pois o canal2 demora mais que o canal1
		// mensagemCanal1 := <-canal1
		// fmt.Println(mensagemCanal1)
		// mensagemCanal2 := <-canal2
		// fmt.Println(mensagemCanal2)

		// Com o uso do select ele ira executar mais rápido o canal 1
		select{
			case mensagemCanal1 := <- canal1:
					fmt.Println(mensagemCanal1)

			case mensagemCanal2 := <- canal2:
					fmt.Println(mensagemCanal2)
		}
	}
}