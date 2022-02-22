package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("sqlite3", config.StringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
