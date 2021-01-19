package router

import (
	"api/src/router/rotas"
	"github.com/gorilla/mux"
)

// Gerar vai retornar um router com as Rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter();
	return rotas.Configurar(r)
}