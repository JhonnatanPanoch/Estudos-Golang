package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarUsuario cria um novo usuário no banco de dados
func CriarUsuario(rw http.ResponseWriter, r *http.Request) {
	request, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(request, &usuario); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.CriarRepositorioUsuarios(db)
	usuario.Id, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(rw, http.StatusCreated, usuario)
}

// BuscarUsuarios obtém uma lista de usuarios do banco de dados pelo seu nome ou nick
func BuscarUsuarios(rw http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.CriarRepositorioUsuarios(db)
	usuarios, erro := repositorio.BuscarPorNomeOuNick(nomeOuNick)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(rw, http.StatusOK, usuarios)
}

// BuscarUsuario obtém um usuario do banco de dados pelo seu id
func BuscarUsuario(rw http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
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

	repositorio := repositorios.CriarRepositorioUsuarios(db)
	usuario, erro := repositorio.BuscarPorId(usuarioId)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(rw, http.StatusOK, usuario)
}

// AtualizarUsuario atualiza um usuario do banco de dados pelo seu id
func AtualizarUsuario(rw http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	idUsuario, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	idUsuarioLogado, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnauthorized, erro)
		return
	}

	if idUsuario != idUsuarioLogado {
		respostas.Erro(rw, http.StatusForbidden, errors.New("não é possível alterar outros usuários"))
		return
	}

	requisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro := json.Unmarshal(requisicao, &usuario); erro != nil {
		respostas.Erro(rw, http.StatusUnprocessableEntity, erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.CriarRepositorioUsuarios(db)
	if erro := repositorio.Alterar(idUsuario, usuario); erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(rw, http.StatusNoContent, nil)
}

// DeletarUsuario apaga um usuario do banco de dados pelo seu id
func DeletarUsuario(rw http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	idUsuario, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	idUsuarioLogado, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnauthorized, erro)
		return
	}

	if idUsuario != idUsuarioLogado {
		respostas.Erro(rw, http.StatusForbidden, errors.New("não é possível apagar outros usuários"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.CriarRepositorioUsuarios(db)
	if erro := repositorio.Excluir(idUsuario); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
	}

	respostas.JSON(rw, http.StatusNoContent, nil)
}

// SeguirUsuario permite que um usuário siga outro
func SeguirUsuario(rw http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	idUsuario, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	seguidorId, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnauthorized, erro)
		return
	}

	if idUsuario == seguidorId {
		respostas.Erro(rw, http.StatusForbidden, errors.New("não é possível seguir a própria conta"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.CriarRepositorioUsuarios(db)
	if erro := repositorio.Seguir(idUsuario, seguidorId); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(rw, http.StatusNoContent, nil)
}

// PararDeSeguirUsuario permite que um usuário deixe de seguir outro
func PararDeSeguirUsuario(rw http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	idUsuario, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	seguidorId, erro := autenticacao.ExtrairUsuarioID(r)
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

	repositorio := repositorios.CriarRepositorioUsuarios(db)
	if erro := repositorio.PararDeSeguir(idUsuario, seguidorId); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(rw, http.StatusNoContent, nil)
}

// BuscarSeguidores efetua a busca dos seguidores do usuário por id
func BuscarSeguidores(rw http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	idUsuario, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
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

	repositorio := repositorios.CriarRepositorioUsuarios(db)
	seguidores, erro := repositorio.BuscarSeguidores(idUsuario)
	if erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(rw, http.StatusOK, seguidores)
}

func BuscarSeguindo(rw http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	idUsuario, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
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

	repositorio := repositorios.CriarRepositorioUsuarios(db)
	seguidores, erro := repositorio.BuscarSeguindo(idUsuario)
	if erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(rw, http.StatusOK, seguidores)
}

// AtualizarSenha permite atualizar a senha de um usuário
func AtualizarSenha(rw http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	idUsuario, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	seguidorId, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnauthorized, erro)
		return
	}

	if idUsuario != seguidorId {
		respostas.Erro(rw, http.StatusForbidden, errors.New("não é possível ataulizar a senha de outro usuário"))
		return
	}

	request, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnprocessableEntity, erro)
		return
	}

	var senha modelos.Senha
	if erro = json.Unmarshal(request, &senha); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.CriarRepositorioUsuarios(db)
	usuarioBanco, erro := repositorio.BuscarPorId(idUsuario)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	if erro := seguranca.VerificarSenha(usuarioBanco.Senha, senha.Atual); erro != nil {
		respostas.Erro(rw, http.StatusUnauthorized, errors.New("a senha atual é inválida"))
		return
	}

	senhaComHash, erro := seguranca.Hash(senha.Nova)
	if erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	if erro := repositorio.AtualizarSenha(idUsuario, string(senhaComHash)); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(rw, http.StatusOK, nil)
}
