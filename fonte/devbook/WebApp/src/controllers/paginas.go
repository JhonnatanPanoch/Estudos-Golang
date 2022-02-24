package controllers

import (
	"net/http"
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
	utils.ExecutarTemplate(rw, "home.html", nil)
}
