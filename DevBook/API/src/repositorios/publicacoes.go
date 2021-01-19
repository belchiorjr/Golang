package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publicacoes representa um repositório de publicações
type Publicacoes struct{
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositório de publicacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes{
	return &Publicacoes{db}
} 

// Criar função que insere uma nova publicação na base de dados
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?,?,?)", 
	)

	if erro != nil{
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID) 

	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}


// BuscarPorID traz uma única publicação do banco de dados
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		SELECT p.*, u.nick 
		FROM publicacoes p 
			INNER JOIN usuarios u 
				ON u.id = p.autor_id 
		WHERE p.id = ?
	`, publicacaoID)

	if erro != nil {
		return modelos.Publicacao{}, erro
	}

	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if erro := linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Buscar traz as publicações dos usuários seguidos e também do próprio usuário que fez a requisição
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT distinct p.*, u.nick FROM publicacoes p
			INNER JOIN usuarios u ON u.id = p.autor_id
			LEFT JOIN seguidores s ON s.usuario_id = p.autor_id
		WHERE u.id = ? OR s.seguidor_id = ?
		ORDER BY 1 DESC
	`, usuarioID, usuarioID)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro := linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Atualizar altera os dados de uma publicacao no banco de dados
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare("UPDATE publicacoes SET titulo = ?, conteudo = ? WHERE id = ?")

	if erro != nil {
		return erro 
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}

	return nil
}


// Deletar exclui uma publicação no banco de dados.
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM publicacoes WHERE id = ?")

	if erro != nil {
		return erro 
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}



// BuscarPorUsuario traz as publicações de um usuário específico
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT p.*, u.nick FROM publicacoes p
		INNER JOIN usuarios u ON u.id = p.autor_id
		WHERE p.autor_id = ? 
		ORDER BY 1 DESC
	`, usuarioID)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro := linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Curtir adiciona 1 valor a mais no campo curtidas
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("UPDATE publicacoes SET curtidas = curtidas + 1 WHERE id = ?")

	if erro != nil {
		return erro 
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}


// Descurtir subtrair 1 valor a menos no campo curtidas
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("UPDATE publicacoes SET curtidas = curtidas - 1 WHERE id = ? AND curtidas >= 0 ")

	if erro != nil {
		return erro 
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}