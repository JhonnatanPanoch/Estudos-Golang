package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

//Usuarios representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// CriarRepositorioUsuarios cria uma instancia de repositório de usuarios
func CriarRepositorioUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar Insere um usuário no banco de dados
func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	var query = `INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)`
	statement, erro := u.db.Prepare(query)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(idInserido), nil
}
