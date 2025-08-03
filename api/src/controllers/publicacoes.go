package controllers

import (
	"api/src/authentication"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositories"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

// CriarPublicacao adiciona uma nova publicação no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioID

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)
}

// BuscarPublicacoes traz as publicações que apareceriam no feed do usuário
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}

// BuscarPublicacao traz uma única publicação
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

}

// AtualizarPublicacao altera os dados de uma publicação
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}

// DeletarPublicacao exlui os dados de uma publicação
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}
