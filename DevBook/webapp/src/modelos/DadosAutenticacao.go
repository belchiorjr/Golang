package modelos

// DadosAutenticacao representa os dados da autenticação no sistema.
type DadosAutenticacao struct {
	ID string `json:"id"`
	Token string `json:"token"`
}