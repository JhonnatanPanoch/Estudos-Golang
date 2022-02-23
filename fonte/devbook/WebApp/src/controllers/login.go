package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// CarregarLogin vai carregar a tela de login
func CarregarLogin(rw http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(rw, "login.html", nil)
}
