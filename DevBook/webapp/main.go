package main

import (
	// "encoding/hex"
	// "github.com/gorilla/securecookie"
	"webapp/src/cookies"
	"webapp/src/config"
	"webapp/src/utils"
	"fmt"
	"webapp/src/router"
	"log"
	"net/http"
)

// func init() {
// 	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(hashKey)
// 	blocKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(blocKey)
// }

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Escutando na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}