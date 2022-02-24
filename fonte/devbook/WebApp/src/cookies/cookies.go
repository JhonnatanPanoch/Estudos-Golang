package cookies

import (
	"fmt"
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Salvar(rw http.ResponseWriter, id, token string) error {

	dados := map[string]string{
		"id":    id,
		"token": token,
	}

	dadosCodificados, erro := s.Encode("dados", dados)
	if erro != nil {
		fmt.Println(erro)
		return erro
	}

	http.SetCookie(rw, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func Ler(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}

	var valores = make(map[string]string)
	if erro = s.Decode("dados", cookie.Value, &valores); erro != nil {
		return nil, erro
	}

	return valores, nil
}
