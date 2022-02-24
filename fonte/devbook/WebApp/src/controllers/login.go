package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/models"
	"webapp/src/respostas"
)

func FazerLogin(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	requisicao, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(rw, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	response, erro := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(requisicao))
	if erro != nil {
		respostas.JSON(rw, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	var dadosAutenticacao models.DadosAutenticacao

	if erro = json.NewDecoder(r.Body).Decode(&dadosAutenticacao); erro != nil {
		respostas.JSON(rw, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
	}

	respostas.JSON(rw, http.StatusOK, nil)
}
