package rotas

import (
	"webapp/src/controllers"
	"net/http"
)

var rotasLogout = []Rota{
	{
		URI: "/logout",
		Metodo:http.MethodGet,
		Funcao:controllers.Logout,
		RequerAutenticacao:true,
	},
}