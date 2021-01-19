package main

import (
	"fmt"
)


func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":"Pedro",
		"sobrenome":"Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])
	
	usuario2 := map[int]string{
		1:"Pedro",
		2:"Silva",
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2[1])
	fmt.Println(usuario2[2])


	usuario3 := map[string]map[string]string{
		"ficha": {
			"nome":"Belchior Pereira de Araújo",
			"idade":"40 anos",
		},
	}

	fmt.Println(usuario3)
	fmt.Println(usuario3["ficha"]["nome"])
	fmt.Println(usuario3["ficha"]["idade"])
	
	delete(usuario3["ficha"],"idade")
	fmt.Println(usuario3)
	
	usuario3["curso"] = map[string]string{
		"Nome do Curso": "Análise de Sistemas",
		"Duração": "4 anos",
	}
	fmt.Println(usuario3)

}