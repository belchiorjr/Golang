package main

import (
	"sync"
	"time"
	"fmt" 
)

func main() { 
	// CONCORRÊNCIA != PARALELISMO
	fmt.Println("Programa Wait Group")

	var waitGroup sync.WaitGroup

	waitGroup.Add(4) // Quantidade de goroutines

	// Utilizar uma função Anônima
	go func() {
		escrever("Go Routine 1")
		waitGroup.Done()
	} ()
		
	go func () {
		escrever("Go Routine 2")
		waitGroup.Done()
	} ()
	
	go func () {
		escrever("Go Routine 3")
		waitGroup.Done()
	} ()
	
	go func () {
		escrever("Go Routine 4")
		waitGroup.Done()
	} ()

	waitGroup.Wait()
	
}


func escrever (texto string) {
	for i:=0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}