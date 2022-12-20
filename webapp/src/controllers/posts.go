package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/response"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	post, erro := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
	}

	url := fmt.Sprintf("%s/posts", config.ApiUrl)

	responseApi, erro := requests.RequestWithAutentication(r, http.MethodPost, url, bytes.NewBuffer(post))
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

	defer responseApi.Body.Close()

	if responseApi.StatusCode >= 400 {
		response.ErrorHandler(w, responseApi)
		return
	}

	response.JSON(w, responseApi.StatusCode, nil)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
	}

	url := fmt.Sprintf("%s/posts/%d/like", config.ApiUrl, postId)

	responseApi, erro := requests.RequestWithAutentication(r, http.MethodPost, url, nil)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

	defer responseApi.Body.Close()

	if responseApi.StatusCode >= 400 {
		response.ErrorHandler(w, responseApi)
		return
	}

	response.JSON(w, responseApi.StatusCode, nil)

}

func DislikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
	}

	url := fmt.Sprintf("%s/posts/%d/dislike", config.ApiUrl, postId)

	responseApi, erro := requests.RequestWithAutentication(r, http.MethodPost, url, nil)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

	defer responseApi.Body.Close()

	if responseApi.StatusCode >= 400 {
		response.ErrorHandler(w, responseApi)
		return
	}

	response.JSON(w, responseApi.StatusCode, nil)

}

func UpdatePost(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	r.ParseForm()

	post, erro := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
	}

	url := fmt.Sprintf("%s/posts/%d", config.ApiUrl, postId)

	responseApi, erro := requests.RequestWithAutentication(r, http.MethodPut, url, bytes.NewBuffer(post))
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

	defer responseApi.Body.Close()

	if responseApi.StatusCode >= 400 {
		response.ErrorHandler(w, responseApi)
		return
	}

	response.JSON(w, responseApi.StatusCode, nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.ApiUrl, postId)

	responseApi, erro := requests.RequestWithAutentication(r, http.MethodDelete, url, nil)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

	response.JSON(w, responseApi.StatusCode, nil)

}
