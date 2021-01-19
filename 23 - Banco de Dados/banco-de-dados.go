package main

import (
	"database/sql"
	"log"
	"fmt" 

	_ "github.com/go-sql-driver/mysql" // utilização implícita para o pacote database/sql
)

func main() { 
	fmt.Println("Programa Banco de Dados Mysql")

	stringConexao := "root:@/golang?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		fmt.Println("Dentro do SQL.Open")
		log.Fatal(erro)
	}
	
	defer db.Close()
	
	if erro = db.Ping(); erro != nil {
		fmt.Println("Dentro do Ping")
		log.Fatal(erro)
	}
	
	fmt.Println("A Conexão com a base de dados está aberta!")
	
	linhas, erro := db.Query("SELECT * FROM usuarios")
	
	if erro = db.Ping(); erro != nil {
		fmt.Println("Dentro do Query")
		log.Fatal(erro)
	}


	defer linhas.Close()

	fmt.Println(linhas)
	
	
	//fmt.Println(db)
}