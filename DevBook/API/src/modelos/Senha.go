package modelos

import (
	"errors"
)

// Senha representa o formato da requisição de alteração de senha
type Senha struct{
	Nova string `json:"nova"`
	Atual string `json:"atual"`
}

func (senha *Senha) validar() error{

	if senha.Nova == "" {
		return errors.New("A Nova Senha é obrigatória e não pode estar em branco")
	}

	if senha.Atual == "" {
		return errors.New("A Senha Atual é obrigatória e não pode estar em branco")
	}

	return nil
}