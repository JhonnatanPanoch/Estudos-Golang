package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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

func BuscarUsuarios(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Buscando usuários"))
}

func BuscarUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Buscando usuário"))
}

func AtualizarUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Atualizando usuário"))
}

func DeletarUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Apagando usuário"))
}
