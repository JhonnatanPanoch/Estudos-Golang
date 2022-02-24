package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"webapp/src/config"
	"webapp/src/models"
	"webapp/src/requisicoes"
	"webapp/src/utils"
)

// CarregarLogin vai carregar a tela de login
func CarregarLogin(rw http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(rw, "login.html", nil)
}

// CarregarPaginaCadastroUsuario Carrega página de cadastro de usuários
func CarregarPaginaCadastroUsuario(rw http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(rw, "cadastro.html", nil)
}

func CarregarPaginaPrincipal(rw http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.ApiUrl)

	resp, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		fmt.Println(erro)
	}

	corpoResposta, erro := ioutil.ReadAll(resp.Body)
	if erro != nil {
		fmt.Println(erro)
	}

	var publicacoes []models.Publicacao
	if erro = json.Unmarshal(corpoResposta, &publicacoes); erro != nil {
		fmt.Println(erro)
	}

	utils.ExecutarTemplate(rw, "home.html", resp)
}
