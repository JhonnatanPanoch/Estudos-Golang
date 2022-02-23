package controllers

import "net/http"

// CarregarLogin vai carregar a tela de login
func CarregarLogin(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("tela de login"))
}
