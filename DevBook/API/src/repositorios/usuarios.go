package repositorios

import (
	"fmt"
	"api/src/modelos"
	"database/sql"
)

// Usuarios representa um repositório de usuários
type Usuarios struct{
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuário
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios{
	return &Usuarios{db}
} 

// Criar função que insere um novo usuário na base de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO usuarios (nome, nick, email, senha) VALUES (?,?,?,?)", 
	)

	if erro != nil{
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha) 

	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}


// Buscar traz todos os usuários que atendem um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? OR nick LIKE ?", 
		nomeOuNick, 
		nomeOuNick,
	)

	if erro != nil{
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Criadoem,
		); erro != nil {
			return nil, erro
		} 

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}


// BuscarPorID traz um usuario do banco de dados
func (repositorio Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ? ", 
		ID,
	)

	if erro != nil{
		return modelos.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Criadoem,
		); erro != nil {
			return modelos.Usuario{}, erro
		} 
	}

	return usuario, nil
}	

// Atualizar dados do usuário
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE usuarios SET nome=?, nick=?, email=? WHERE id = ?", 
	)

	if erro != nil{
		return erro
	}

	defer statement.Close()
	
	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}


// Deletar os dados do usuário no banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}


// BuscarPorEmail obtém os dados do usuário po email
func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT id, senha FROM usuarios WHERE email = ?", 
		email,
	)

	if erro != nil{
		return modelos.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Senha,
		); erro != nil {
			return modelos.Usuario{}, erro
		} 
	}

	return usuario, nil
}


// Seguir relaciona o usuário logado com outro usuario na tabela de seguidores
func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {

	statement, erro := repositorio.db.Prepare(
		"INSERT IGNORE INTO seguidores(usuario_id, seguidor_id) VALUES (?, ?)",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}


// PararDeSeguirUsuario desrelaciona o usuário logado com outro usuario na tabela de seguidores
func (repositorio Usuarios) PararDeSeguirUsuario(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"DELETE FROM seguidores WHERE usuario_id = ? AND seguidor_id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}


// BuscarSeguidores - Traz todos os seguidores de um usuário
func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		`SELECT u.id, u.nome, u.nick, u.email, u.criadoEm FROM usuarios u 
		INNER JOIN seguidores s ON u.id = s.seguidor_id WHERE s.usuario_id  = ?`,
		usuarioID,
	)

	if erro != nil{
		return nil, erro
	}

	defer linhas.Close()

	var seguidores []modelos.Usuario

	for linhas.Next() {
		var seguidor modelos.Usuario

		if erro = linhas.Scan(
			&seguidor.ID,
			&seguidor.Nome,
			&seguidor.Nick,
			&seguidor.Email,
			&seguidor.Criadoem,
		); erro != nil {
			return nil, erro
		} 

		seguidores = append(seguidores, seguidor)
	}

	return seguidores, nil
}


// BuscarSeguindo - Traz todos os usuários que está seguindo
func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		`SELECT u.id, u.nome, u.nick, u.email, u.criadoEm FROM usuarios u 
		INNER JOIN seguidores s ON u.id = s.usuario_id WHERE s.seguidor_id = ?`,
		usuarioID,
	)

	if erro != nil{
		return nil, erro
	}

	defer linhas.Close()

	var seguidores []modelos.Usuario

	for linhas.Next() {
		var seguidor modelos.Usuario

		if erro = linhas.Scan(
			&seguidor.ID,
			&seguidor.Nome,
			&seguidor.Nick,
			&seguidor.Email,
			&seguidor.Criadoem,
		); erro != nil {
			return nil, erro
		} 

		seguidores = append(seguidores, seguidor)
	}

	return seguidores, nil
}


// BuscarSenha - Traz a senha do usuário por id
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linhas, erro := repositorio.db.Query("SELECT senha FROM usuarios WHERE id = ?", usuarioID)

	if erro != nil{
		return "", erro
	}

	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(&usuario.Senha); erro != nil {
			return "", erro
		} 
	}

	return usuario.Senha, nil
}


// AtualizarSenha - Altera a senha de um usuário
func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE usuarios SET senha=? WHERE id = ?", 
	)

	if erro != nil{
		return erro
	}

	defer statement.Close()
	
	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}

	return nil
}

