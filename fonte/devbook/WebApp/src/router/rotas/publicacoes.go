package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasPublicacao = []Rota{
	{
		URI:                "publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
}
