package main

import (
	"time"
	"fmt" 
)

func main() { 
	// CONCORRÊNCIA != PARALELISMO
	fmt.Println("Programa Goroutines")

	go escrever("Olá Mundo!")
	escrever("Programando em Go")
}


func escrever (texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}