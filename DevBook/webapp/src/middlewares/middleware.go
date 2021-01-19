package middlewares

import (
	//"fmt"
	"webapp/src/cookies"
	"log"
	"net/http"
)

// Logger escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Autenticar verifica a existência de cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)

		if _, erro := cookies.Ler(r); erro != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}

		//fmt.Println(valores, erro)
		proximaFuncao(w, r)
	}
}


