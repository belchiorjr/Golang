package main

import (
	"fmt" 
)

func escrever() {
	fmt.Println("Escrevendo...")
}

type usuario struct {
	nome string
	idade uint8
}

func(u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18	
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("Métodos")
	escrever()
	
	usuario1 := usuario{"Belchior Junior", 20}
	fmt.Println(usuario1)
	fmt.Println(usuario1.maiorDeIdade())
	usuario1.salvar()
	
	usuario2 := usuario{"Taiane Lopes de Lima", 10}
	fmt.Println(usuario2)
	fmt.Println(usuario2.maiorDeIdade())
	fmt.Println(usuario2.idade)

	usuario2.fazerAniversario()

	fmt.Println(usuario2.idade)
	usuario2.salvar()
}