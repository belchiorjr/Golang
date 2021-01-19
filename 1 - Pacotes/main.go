package main

import (
	"fmt"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo no arquivo main")
	auxiliar.Escrever()
	auxiliar.Escrever()

	// email, err := emailaddress.Parse("foobar.com")

	// if err != nil {
	// 		fmt.Println("Email inválido!")
	// }

	// fmt.Println(email.LocalPart) // foo
	// fmt.Println(email.Domain) // bar.com
	// fmt.Println(email) // foo@bar.com
	// fmt.Println(email.String()) 

}