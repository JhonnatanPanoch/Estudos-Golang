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

	linhas, erro := p.db.Query("SELECT * FROM publicacoes WHERE id = ?", publicacaoId)
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
		); erro != nil {
			return publicacao, erro
		}
	}

	return publicacao, nil
}
