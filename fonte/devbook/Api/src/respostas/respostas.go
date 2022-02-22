package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma resposta em JSON para a requisição.
func JSON(rw http.ResponseWriter, statusCode int, dados interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)

	if erro := json.NewEncoder(rw).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

func Erro(rw http.ResponseWriter, statusCode int, erro error) {

	strErro := struct {
		Erro string `json:"erro"`
		Code int    `json:"code"`
	}{
		Erro: erro.Error(),
		Code: statusCode,
	}

	JSON(rw, statusCode, strErro)
}
