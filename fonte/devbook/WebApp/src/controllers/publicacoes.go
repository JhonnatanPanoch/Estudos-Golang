package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

func CriarPublicacao(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publicacao, erro := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})

	if erro != nil {
		respostas.JSON(rw, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes", config.ApiUrl)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))
	if erro != nil {
		respostas.JSON(rw, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(rw, response)
		return
	}

	respostas.JSON(rw, response.StatusCode, nil)
}

func CurtirPublicacao(rw http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(rw, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d/curtir", config.ApiUrl, publicacaoId)
	resp, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(rw, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respostas.JSON(rw, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	respostas.JSON(rw, http.StatusOK, nil)
}

func DescurtirPublicacao(rw http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(rw, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d/descurtir", config.ApiUrl, publicacaoId)
	resp, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(rw, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respostas.JSON(rw, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	respostas.JSON(rw, http.StatusOK, nil)
}
