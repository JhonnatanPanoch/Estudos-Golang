package midlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"net/http"
)

// Autenticar verifica se o usuário que está fazendo a requisção está autenticado
// HandlerFunc = rw http.ResponseWriter, r *http.Request
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(rw, http.StatusUnauthorized, erro)
			return
		}
		next(rw, r)

	}
}
