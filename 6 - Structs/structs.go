package main

import (
	"fmt"
)

type usuario struct{
	nome string
	idade uint8
	endereco endereco
}

type endereco struct{
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var user usuario
	fmt.Println(user)
	
	user.nome = "Belchior"
	user.idade = 40
	fmt.Println(user)

	user2 := usuario{nome:"Davi", idade:21}
	fmt.Println(user2)

	user3 := usuario{nome: "Davi"}
	fmt.Println(user3)

	enderecoExemplo := endereco{"Travessa Dário Veloso", 41}

	user4 := usuario{nome:"Davi",  endereco: enderecoExemplo}
	fmt.Println(user4)
	
}