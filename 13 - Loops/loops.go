package main

import (
	"time"
	"fmt"
)


func main() {
	fmt.Println("Loops")

	// i := 0
	
	// for i < 10 {a
	// 	i++
	// 	fmt.Print(" Incrementando i ")
	// 	time.Sleep(time.Second)
	// }
	
	// fmt.Println(i)
	
	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j ", j)	
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string {"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}


	for indice, letra  := range "PALAVRA" {
		//fmt.Println(indice, letra)
		fmt.Println(indice, letra, string(letra))
	}

	usuario := map[string]string{
		"nome": "Leonardo",
		"sobrenome":"Silva",
	}

	fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	type usuarioStruct struct {
		nome string
		sobrenome string
	}


	// usuario2 := usuarioStruct{"José", "Júnior"}

	// for chave, valor := range usuario2 {
	// 	fmt.Println(chave, valor)
	// }


	// Loooping infinito

	for {
		fmt.Println("Executando inifinitamente!!!!")
		time.Sleep(time.Second)
	}

}