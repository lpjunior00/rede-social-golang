package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {

	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro := json.Unmarshal(bodyRequest, &user); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	userRepository := repositories.NewUserRepository(db)
	userDB, erro := userRepository.FindByEmail(user.Email)
	if erro != nil {
		response.Erro(w, http.StatusNotFound, erro)
		return
	}

	if erro := security.ComparePassword(userDB.Password, user.Password); erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := auth.CreateToken(userDB.ID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	userId := strconv.FormatUint(userDB.ID, 10)

	response.JSON(w, http.StatusOK, models.AuthData{UserId: userId, Token: token})

}
