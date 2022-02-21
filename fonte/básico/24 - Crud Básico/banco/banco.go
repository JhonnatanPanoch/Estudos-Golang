package banco

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "./devbook.db"

	db, erro := sql.Open("sqlite3", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
