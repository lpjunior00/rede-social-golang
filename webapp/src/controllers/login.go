package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/response"
)

func Login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	user, erro := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.ApiUrl)
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

	var authData models.AuthData
	if erro = json.NewDecoder(responseBody.Body).Decode(&authData); erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErroAPI{Erro: erro.Error()})
		return
	}

	if erro = cookies.SaveCookie(w, authData.UserId, authData.Token); erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

}
