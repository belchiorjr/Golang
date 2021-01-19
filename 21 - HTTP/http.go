package main

import (
	"net/http"
	"log"
	"fmt" 
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuários!"))
}

func raiz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página Raiz!"))
}

func main() {
	fmt.Println("Programa HTTP")
	fmt.Println("HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB")
	fmt.Println("ARQUITETURA CLIENTE-SERVIDOR")
	fmt.Println("CLIENTE FAZ A REQUISIÇÃO, SERVIDOR PROCESSA E ENVIA A RESPOSTA")
	fmt.Println("REQUEST - RESPONSE")
	fmt.Println("ROTAS - URI: Identificador do Recurso")
	fmt.Println("ROTAS - Métodos: GET, POST, PUT, DELETE")

	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)
	http.HandleFunc("/", raiz)

	log.Fatal(http.ListenAndServe(":5000", nil))
}