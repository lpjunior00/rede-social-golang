package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requests"
	"webapp/src/response"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	user, erro := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"nickname": r.FormValue("nickname"),
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	fmt.Println(bytes.NewBuffer(user))

	url := fmt.Sprintf("%s/users", config.ApiUrl)
	responseBody, erro := http.Post(url, "application/json", bytes.NewBuffer(user))

	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

	defer responseBody.Body.Close()

	if responseBody.StatusCode >= 400 {
		response.ErrorHandler(w, responseBody)
		return
	}

	response.JSON(w, responseBody.StatusCode, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userId, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/unfollow/%d", config.ApiUrl, userId)
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

	response.JSON(w, http.StatusOK, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userId, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/follow/%d", config.ApiUrl, userId)
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

	response.JSON(w, http.StatusOK, nil)
}

func EditUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	user, erro := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nickname": r.FormValue("nickname"),
	})

	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.ReadCookie(r)
	loggerUserId, _ := strconv.ParseUint(cookie["userId"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, loggerUserId)
	responseApi, erro := requests.RequestWithAutentication(r, http.MethodPut, url, bytes.NewBuffer(user))

	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

	defer responseApi.Body.Close()
	if responseApi.StatusCode >= 400 {
		response.ErrorHandler(w, responseApi)
		return
	}

	response.JSON(w, http.StatusOK, nil)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	passwords, erro := json.Marshal(map[string]string{
		"currentPassword": r.FormValue("currentPassword"),
		"newPassword":     r.FormValue("newPassword"),
	})

	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.ReadCookie(r)
	userId, _ := strconv.ParseUint(cookie["userId"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/update-password", config.ApiUrl, userId)
	responseApi, erro := requests.RequestWithAutentication(r, http.MethodPost, url, bytes.NewBuffer(passwords))
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

	defer responseApi.Body.Close()

	if responseApi.StatusCode >= 400 {
		response.ErrorHandler(w, responseApi)
		return
	}

	response.JSON(w, http.StatusOK, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookie(r)
	userId, _ := strconv.ParseUint(cookie["userId"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, userId)
	responseApi, erro := requests.RequestWithAutentication(r, http.MethodDelete, url, nil)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

	defer responseApi.Body.Close()

	if responseApi.StatusCode >= 400 {
		response.ErrorHandler(w, responseApi)
		return
	}

	response.JSON(w, http.StatusOK, nil)
}
