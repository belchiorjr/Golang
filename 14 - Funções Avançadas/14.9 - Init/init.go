
package main

import (
	"fmt" 
)

var n int

func init () {
	fmt.Println("Executando a função init")
	n = 1000
}

func main() { 
	fmt.Println("Função init")
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}