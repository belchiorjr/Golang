package main

import (
	"log"
	"encoding/json"
	"fmt" 
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string	`json:"raca"`
	Idade uint	`json:"idade"`
}

func main() { 
	fmt.Println("Programa Unmarshal")
	cachorroEmJSON := `{"nome":"Rex","raca":"Dalmata","idade":3}`
	fmt.Println(cachorroEmJSON)

	//c := cachorro{}
	var c cachorro
	
	// Precisa tratar os erros
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)
	
	// Transformar em map
	cachorro2EmJSON := 	`{"nome":"Toby","raca":"Poodle"}`
	
	c2 := make(map[string]string)
	
	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil{
		log.Fatal(erro)
	}
	
	fmt.Println(c2)
}