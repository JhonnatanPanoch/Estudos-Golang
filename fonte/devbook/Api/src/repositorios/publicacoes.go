package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publicacoes representa um repositorio de publicações
type Publicacoes struct {
	db *sql.DB
}

func CriarRepositorioPublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (p Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	var query = `INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)`
	statement, erro := p.db.Prepare(query)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	insercao, erro := statement.Exec(
		publicacao.Titulo,
		publicacao.Conteudo,
		publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(idInserido), nil
}

func (p Publicacoes) BuscarPorId(publicacaoId uint64) (modelos.Publicacao, error) {
	var publicacao = modelos.Publicacao{}

	linhas, erro := p.db.Query(`
	select p.*, u.nick from 
	publicacoes p inner join usuarios u
	on u.id = p.autor_id where p.id = ?`, publicacaoId)
	if erro != nil {
		return publicacao, erro
	}
	defer linhas.Close()

	if linhas.Next() {
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return publicacao, erro
		}
	}

	return publicacao, nil
}

func (p Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {

	linhas, erro := p.db.Query(`
		select distinct p.*, u.nick from publicacoes p 
		inner join usuarios u on u.id = p.autor_id 
		inner join seguidores s on p.autor_id = s.usuario_id 
		where u.id = ? or s.seguidor_id = ?
		order by 1 desc;`, usuarioID, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao = modelos.Publicacao{}
		if erro = linhas.Scan(
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