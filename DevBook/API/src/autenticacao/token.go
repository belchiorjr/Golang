package autenticacao

import (
	"strconv"
	"errors"
	"fmt"
	"strings"
	"net/http"
	"api/src/config"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken - Define o token de segurança
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString([]byte(config.SecretKey)) // secret
}

// ValidarToken verifica se o token passado na requisição é válido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)

	if erro != nil {
		return erro
	}
	
	// Aqui é feita a validção do token, verifica se foi encontrado os Claims e se o token é valido
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil // Retorna nil pois está tudo certo, para prosseguir
	}

	return errors.New("Token Inválido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatuana inesperado! %v", token.Header["alg"])
	}
	
	return config.SecretKey, nil
}

// ExtrairUsuarioID - Obtém o id do usuário 
func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)

	if erro != nil {
		return 0, erro
	}
	
	// Aqui é feita a validção do token, verifica se foi encontrado os Claims e se o token é valido
	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]) , 10, 64)
		
		if erro != nil {
			return 0, erro
		}

		return usuarioID, nil
	}

	return 0, errors.New("Token Inválido")
}