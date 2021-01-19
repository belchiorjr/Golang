package rotas


import (
	"webapp/src/controllers"
	"net/http"
)

var rotasPaginaPrincipal = []Rota{
	{
		URI: "/home",
		Metodo:http.MethodGet,
		Funcao:controllers.CarregarPaginaPrincipal,
		RequerAutenticacao: true,
	},
}