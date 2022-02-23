package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.CarregarTemplates()
	rotas := router.Gerar()

	fmt.Println("Rodando a aplicação web na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", rotas))
}
