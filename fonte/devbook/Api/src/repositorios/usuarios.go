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

// BuscarPorNomeOuNick traz todos os usuarios que atendam um filtro de nome ou nick
func (u Usuarios) BuscarPorNomeOuNick(nomeOuNick string) ([]modelos.Usuario, error) {
	linhas, erro := u.db.Query("select * from usuarios where nome = ? or nick = ?", nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro := linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Senha,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorId traz um usuário do banco de dados pelo id.
func (u Usuarios) BuscarPorId(id uint64) (modelos.Usuario, error) {
	var usuario = modelos.Usuario{}

	linhas, erro := u.db.Query("select * from usuarios where id = ?", id)
	if erro != nil {
		return usuario, erro
	}
	defer linhas.Close()

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Senha,
			&usuario.CriadoEm,
		); erro != nil {
			return usuario, erro
		}
	}

	return usuario, erro
}

// Alterar altera as informaçoes básicas de um usuário do banco de dados pelo seu id
func (u Usuarios) Alterar(idUsuario uint64, usuario modelos.Usuario) error {
	var query = `UPDATE usuarios SET
					nome = ?, 
					nick = ?, 
					email = ?
				WHERE id = ?;`
	statement, erro := u.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, idUsuario)
	if erro != nil {
		return erro
	}

	return nil
}

func (u Usuarios) Excluir(idUsuario uint64) error {
	statement, erro := u.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(idUsuario)
	if erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail traz o usuario que atendam um filtro de email
func (u Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	var usuario = modelos.Usuario{}
	linhas, erro := u.db.Query("select * from usuarios where email = ?", email)
	if erro != nil {
		return usuario, erro
	}
	defer linhas.Close()

	if linhas.Next() {
		if erro := linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Senha,
			&usuario.CriadoEm,
		); erro != nil {
			return usuario, erro
		}
	}

	return usuario, nil
}

// Seguir permite que dois usuários se sigam
func (u Usuarios) Seguir(idUsuario, idSeguidor uint64) error {

	var query = `INSERT IGNORE INTO seguidores (usuario_id, seguidor_id) VALUES (?, ?)`
	statement, erro := u.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(idUsuario, idSeguidor)
	if erro != nil {
		return erro
	}

	return nil
}
