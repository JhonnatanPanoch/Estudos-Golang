package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	fmt.Println("Rodando a aplicação web na porta 3000")

	rotas := router.Gerar()

	log.Fatal(http.ListenAndServe(":3000", rotas))
}
