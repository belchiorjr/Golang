package main

import (
	"fmt" 
)


func generica (interf interface{}) {
	fmt.Println(interf)
}

func main() { 
	fmt.Println("Interfaces do Tipo Genérico")

	generica("String")
	generica(10)
	generica(100.5)

	mapaGambiarra := map[interface{}]interface{} {
		1:"String",
		float32(100) : true,
		"String" : int(100000),
		true: float64(4654.4566),
	}

	fmt.Println(mapaGambiarra)

}