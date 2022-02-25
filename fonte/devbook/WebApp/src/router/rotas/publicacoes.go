package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

// NÃO ESQUECER DA / NA FRENTE
var rotasPublicacao = []Rota{
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},

	{
		URI:                "/publicacoes/{publicacaoId}/curtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CurtirPublicacao,
		RequerAutenticacao: true,
	},

	{
		URI:                "/publicacoes/{publicacaoId}/descurtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DescurtirPublicacao,
		RequerAutenticacao: true,
	},
}
