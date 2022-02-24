package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ApiUrl   = ""
	AppPort  = 0
	HashKey  []byte
	BlockKey []byte
)

// Carregar vai carregar as variaveis de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	AppPort, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		AppPort = 3000
	}

	ApiUrl = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
