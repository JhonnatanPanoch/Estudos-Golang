package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da aplicação
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(rw http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

// Configurar aplica configuração das rotas das telas
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuario...)

	for _, rota := range rotas {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	fileServer := http.FileServer(http.Dir("./src/assets/"))
	router.PathPrefix("/src/assets/").Handler(http.StripPrefix("/src/assets/", fileServer))

	return router
}
