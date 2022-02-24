package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
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
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respostas.JSON(rw, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	var publicacoes []models.Publicacao
	if erro = json.NewDecoder(resp.Body).Decode(&publicacoes); erro != nil {
		respostas.JSON(rw, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	var dados = struct {
		Publicacoes []models.Publicacao
		UsuarioID   uint64
	}{
		Publicacoes: publicacoes,
		UsuarioID:   usuarioId,
	}

	utils.ExecutarTemplate(rw, "home.html", dados)
}
