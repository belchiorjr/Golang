package controllers

import (
	"webapp/src/cookies"
	"fmt"
	"webapp/src/config"
	"webapp/src/modelos"
	"webapp/src/respostas.go"
	"bytes"
	"encoding/json"
	"net/http"
)

// FazerLogin Utiliza o e-mail e senha do usuário para autenticar na aplicação
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	
	//fmt.Println(usuario)
	//fmt.Println(bytes.NewBuffer(usuario))

	url := fmt.Sprintf("%s/login", config.APIURL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	// token, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(response.StatusCode, string(token))

	// Se ocorrer erro lá na api
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var dadosAutenticacao modelos.DadosAutenticacao

	if erro = json.NewDecoder(response.Body).Decode(&dadosAutenticacao); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if erro = cookies.Salvar(w, dadosAutenticacao.ID, dadosAutenticacao.Token); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}


	respostas.JSON(w, http.StatusOK, nil)
}