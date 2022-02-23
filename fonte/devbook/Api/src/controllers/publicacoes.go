package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CriarPublicao efetua a criação de uma nova publicação
func CriarPublicao(rw http.ResponseWriter, r *http.Request) {
	usuarioId, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioId

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	db.Close()

	repositorio := repositorios.CriarRepositorioPublicacoes(db)
	idPublicacao, erro := repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(rw, http.StatusCreated, idPublicacao)
}

// BuscarPublicacacoes traz as publicações que aparecem no feed do usuário
func BuscarPublicacacoes(rw http.ResponseWriter, r *http.Request) {

}

// BuscarPublicacao traz uma única publicação
func BuscarPublicacao(rw http.ResponseWriter, r *http.Request) {

}

// AtualizarPublicacao efetua a atualização de uma publicação
func AtualizarPublicacao(rw http.ResponseWriter, r *http.Request) {

}

// ApagarPublicacao apaga uma publicação do usuário logado
func ApagarPublicacao(rw http.ResponseWriter, r *http.Request) {

}
