COMANDOS CLI
go build 					- compilar
go mod init modulo 			- criar novo modulo
go install 					- compila arquivo na raiz do go
go get caminhopacote    		- importa novo pacote externo
go mod tidy 				- remove todas as dependencias sem uso
go test					- Efetua o teste do método do arquivo _test
go test ./...  				- Efetua o teste do método de TODOS os arquivo _test
go test --cover				- Efetua os testes e mostra a porcentagem de cobertura
go test --coverprofile file.txt 	- Efetua os testes e exporta um arquivo codificado no diretório
go tool cover --func=file.txt		- Traz informações detalhadas do arquivo gerado a partir do comando acima no console
go tool cover --html=file.txt 	- Traz informações detalhadas do arquivo gerado a partir do comando acima em html

LEMBRETES
- Não é orientada a objeto
- Funções que começam com letra maiúscula são públicas, com letras minúsculas privadas.
- Quando a função ou variável tem * na frente, é ponteiro. Para passar um ponteiro ao invez de um valor usa-se &


Canais
 - (canal <-) recebe um dado e (<- canal) envia um dado
 - Deadlock é quando voce nao tem nenhum lugar enviando dado pro canal mas ele ainda está recebendo dados
 - Para evitar deadlock, fechar o canal e verificar se está fechado
 - Para enviar e receber dados nos canais dentro de uma mesma funçao, utilizar canais com buffer


Testes automatizados
 - Go já tem um pacote para teste já incluso nativamente chamado testing
 - Para o go reconhecer um paocote de testes, precisa ter o _test no nome do arquivo .go
 - A função de teste precisa começar com TestXxxxxXxxx e precisa ter a assintatura (t *testing.T)
 - Para arquivos de testes é possivel ter mais de um arquivo por pacote utilizando _test no nome. Importar o arquivo principal usando ."pasta/arquivo"

Banco de dados
 - Instalação SQLITE no site: https://sqlitebrowser.org/
 - Driver sqlite3: go get https://github.com/mattn/go-sqlite3



instalações 
- mux: 	go get github.com/gorilla/mux
- sqlite3: 	go get github.com/mattn/go-sqlite3
- godotenv  go get github.com/joho/godotenv