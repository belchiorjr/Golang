package banco

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // utilização implícita para o pacote database/sql
)

// Conectar - Realiza a conexão com a base de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "root:@/golang?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		return nil, erro
	}
	
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	
	return db, nil
}