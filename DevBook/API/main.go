package main

import (
	// "encoding/base64"
	// "crypto/rand"
	"api/src/config"
	"log"
	"net/http"
	"api/src/router"
	"fmt" 
)

// func init() {
// 	chave := make([]byte, 64)
	
// 	if _, erro := rand.Read(chave); erro != nil {
// 		log.Fatal(erro)
// 	}
	
// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)

// 	fmt.Println(stringBase64)
// }


func main() { 
	config.Carregar()
	//fmt.Println(config.SecretKey)

	fmt.Printf("Rodando API, escutando na porta: %d", config.Porta)

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}