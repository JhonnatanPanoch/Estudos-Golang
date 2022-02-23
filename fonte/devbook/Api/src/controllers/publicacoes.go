package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioId

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

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
	idUsuario, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.CriarRepositorioPublicacoes(db)
	publicacoes, erro := repositorio.Buscar(idUsuario)
	if erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(rw, http.StatusOK, publicacoes)
}

// BuscarPublicacao traz uma única publicação
func BuscarPublicacao(rw http.ResponseWriter, r *http.Request) {
	var parametros = mux.Vars(r)

	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.CriarRepositorioPublicacoes(db)
	publicacao, erro := repositorio.BuscarPorId(publicacaoId)
	if erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(rw, http.StatusOK, publicacao)
}

// AtualizarPublicacao efetua a atualização de uma publicação
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.CriarRepositorioPublicacoes(db)
	publicacaoSalvaNoBanco, erro := repositorio.BuscarPorId(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar uma publicação que não seja sua"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.Atualizar(publicacaoID, publicacao); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// ApagarPublicacao apaga uma publicação do usuário logado
func ApagarPublicacao(rw http.ResponseWriter, r *http.Request) {

}
