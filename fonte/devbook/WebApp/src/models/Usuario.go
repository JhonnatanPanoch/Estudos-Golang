package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

type Usuario struct {
	Id          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

func BuscarUsuarioCompleto(usuarioId uint64, r *http.Request) (Usuario, error) {

	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarusuarioCarregado(canalUsuario, usuarioId, r)
	go BuscarDadosSeguidores(canalSeguidores, usuarioId, r)
	go BuscarDadosSeguindo(canalSeguindo, usuarioId, r)
	go BuscarDadosPublicacao(canalPublicacoes, usuarioId, r)

	var (
		usuario     Usuario
		seguidores  []Usuario
		seguindo    []Usuario
		publicacoes []Publicacao
	)

	// Caso a rotina chegue até o final desse for,
	// isso significa que as 4 requisições foram chamadas
	// e dados foram retornados com sucesso.
	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.Id == 0 {
				return Usuario{}, errors.New("erro ao buscar dados do usuário")
			}
			usuario = usuarioCarregado

		case seguidoresCarregado := <-canalSeguidores:
			if seguidoresCarregado == nil {
				return Usuario{}, errors.New("erro ao carregar os seguidores do usuário")
			}
			seguidores = seguidoresCarregado

		case seguindoCarregado := <-canalSeguindo:
			if seguindoCarregado == nil {
				return Usuario{}, errors.New("erro ao carregar os usuários que o usuário está seguindo")
			}
			seguindo = seguindoCarregado

		case publicacoesCarregadas := <-canalPublicacoes:
			if publicacoesCarregadas == nil {
				return Usuario{}, errors.New("erro ao carregar as publicações do usuário")
			}

			publicacoes = publicacoesCarregadas
		}

	}

	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacoes

	return usuario, nil

}

func BuscarusuarioCarregado(canal chan<- Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d", config.ApiUrl, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- Usuario{}
		return
	}
	defer response.Body.Close()

	var usuario Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		canal <- Usuario{}
		return
	}

	canal <- usuario
}

func BuscarDadosSeguindo(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.ApiUrl, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var usuarios []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuarios); erro != nil {
		canal <- nil
		return
	}

	canal <- usuarios
}

func BuscarDadosSeguidores(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.ApiUrl, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var usuarios []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuarios); erro != nil {
		canal <- nil
		return
	}

	canal <- usuarios
}

func BuscarDadosPublicacao(canal chan<- []Publicacao, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.ApiUrl, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var publicacoes []Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canal <- nil
		return
	}

	canal <- publicacoes
}
