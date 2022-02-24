package middlewares

import (
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
		if _, erro := cookies.Ler(r); erro != nil {
			http.Redirect(rw, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		next(rw, r)
	}
}
