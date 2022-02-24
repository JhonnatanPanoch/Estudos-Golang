package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	cookies.Configurar()

	utils.CarregarTemplates()
	rotas := router.Gerar()

	fmt.Printf("Rodando a aplicação web na porta :%d", config.AppPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.AppPort), rotas))
}
