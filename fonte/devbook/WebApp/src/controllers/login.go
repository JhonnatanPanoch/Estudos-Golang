package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
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

	url := fmt.Sprintf("%s/login", config.ApiUrl)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(requisicao))
	if erro != nil {
		respostas.JSON(rw, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(rw, response)
		return
	}

	var dadosAutenticacao models.DadosAutenticacao

	if erro = json.NewDecoder(response.Body).Decode(&dadosAutenticacao); erro != nil {
		respostas.JSON(rw, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if erro = cookies.Salvar(rw, dadosAutenticacao.Id, dadosAutenticacao.Token); erro != nil {
		fmt.Println(erro)
		respostas.JSON(rw, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	respostas.JSON(rw, http.StatusOK, nil)
}

func FazerLogout(rw http.ResponseWriter, r *http.Request) {
	cookies.Apagar(rw)
	CarregarLogin(rw, r)

	//http.Redirect(rw, r, "/login", http.StatusTemporaryRedirect)
}
