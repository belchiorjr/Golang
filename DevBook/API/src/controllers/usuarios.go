package controllers

import (
	"api/src/seguranca"
	"errors"
	"api/src/autenticacao"
	"github.com/gorilla/mux"
	"strconv"
	"strings"
	"api/src/respostas"
	"api/src/repositorios"
	"api/src/banco"
	"encoding/json"
	"io/ioutil"
	"api/src/modelos"
	"net/http"
)

// CriarUsuario - Recebe os dados do usuário e repassa para o repositório gravar
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario

	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)
}


// BuscarUsuarios - Lista  os dados dos usuários na base de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	
	db, erro := banco.Conectar()
	
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	
	defer db.Close()
	
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	
	respostas.JSON(w, http.StatusOK, usuarios)
}

// BuscarUsuario - Bucar os dados do usuário na base de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscarPorID(usuarioID)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	
	respostas.JSON(w, http.StatusOK, usuario)
}

// AtualizarUsuario - AtualizarUsuario os dados do usuário na base de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)

	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	
	if usuarioID != usuarioIDNoToken {
		erro := errors.New("Não é permitido alterar dados de um outro usuário")
		respostas.Erro(w, http.StatusForbidden, erro)
		return
	}


	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	
	var usuario modelos.Usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	
	respostas.JSON(w, http.StatusNoContent, nil)

}

// DeletarUsuario - Remove os dados do usuário na base de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)

	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	
	if usuarioID != usuarioIDNoToken {
		erro := errors.New("Não é permitido excluir dados de um outro usuário")
		respostas.Erro(w, http.StatusForbidden, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	if erro = repositorio.Deletar(usuarioID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	
	respostas.JSON(w, http.StatusNoContent, nil)
}


// SeguirUsuario - Permite que um usuário siga outro
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)

	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é permitido seguir a si mesmo"))
		return
	}


	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	if erro = repositorio.Seguir(usuarioID, seguidorID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	
	respostas.JSON(w, http.StatusNoContent, nil)
}



// PararDeSeguirUsuario - Permite que um usuário deixe de seguir outro
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)

	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é permitido deixar de seguir a si mesmo"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	if erro = repositorio.PararDeSeguirUsuario(usuarioID, seguidorID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	
	respostas.JSON(w, http.StatusNoContent, nil)
}


// BuscarSeguidores - Traz todos os seguidores de um usuário
func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {	
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	seguidores, erro := repositorio.BuscarSeguidores(usuarioID)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	
	respostas.JSON(w, http.StatusOK, seguidores)
}



// BuscarSeguindo - Traz todos os usuários que está seguindo
func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {	
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	seguidores, erro := repositorio.BuscarSeguindo(usuarioID)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	
	respostas.JSON(w, http.StatusOK, seguidores)
}


// AtualizarSenha - Permite alterar a senha de um usuário
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)

	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if usuarioID != usuarioIDNoToken {
		erro := errors.New("Não é permitido alterar a senha de um outro usuário")
		respostas.Erro(w, http.StatusForbidden, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	var senha modelos.Senha;

	if erro = json.Unmarshal(corpoRequisicao, &senha); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	senhaSalvaNoBanco, erro := repositorio.BuscarSenha(usuarioID)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	
	if erro = seguranca.VerificarSenha(senhaSalvaNoBanco, senha.Atual); erro != nil {
		erro := errors.New("A Senha salva não condiz com a senha informada")
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	senhaComHash, erro := seguranca.Hash(senha.Nova)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.AtualizarSenha(usuarioID, string(senhaComHash)); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}
