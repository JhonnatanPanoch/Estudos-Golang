package middlewares

import (
	"fmt"
	"net/http"
	"webapp/src/cookies"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		//fmt.Println("\n", r.Method, r.Host, r.RequestURI)
		next(rw, r)
	}
}

func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		dados, erro := cookies.Ler(r)
		if erro != nil {
			fmt.Println(erro)
			http.Redirect(rw, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		fmt.Println(dados)
		next(rw, r)
	}
}
