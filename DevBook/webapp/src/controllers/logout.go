package controllers

import (
	"webapp/src/cookies"
	"net/http"
)

// Logout remove os dados de autenticação salvos no browser
func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", 302)
}