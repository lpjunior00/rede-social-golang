package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	//Read body content
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	//transform body into a user entity
	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//Validate and fix data from user entity
	if erro := user.Prepare("create"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//Connect database
	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	//Create a user repository and save user entity
	userRepository := repositories.NewUserRepository(db)
	userId, erro := userRepository.Create(user)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//return id created
	user.ID = userId
	response.JSON(w, http.StatusCreated, user)
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNickName := strings.ToLower(r.URL.Query().Get("user"))

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	users, erro := userRepository.Find(nameOrNickName)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, users)
}

func FindSpecificUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userId, erro := strconv.ParseUint(parametros["userId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	user, erro := userRepository.FindById(userId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, user)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	userIdInToken, erro := auth.ExtractUserId(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userId != userIdInToken {
		response.Erro(w, http.StatusUnauthorized, errors.New("user cannot update another user"))
		return
	}

	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	//transform body into a user entity
	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//Validate and fix data from user entity
	if erro := user.Prepare("update"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//Connect database
	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	//Create a user repository and save user entity
	userRepository := repositories.NewUserRepository(db)
	if _, erro := userRepository.UpdateUser(user, userId); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//return user updated
	response.JSON(w, http.StatusCreated, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userId, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	userIdInToken, erro := auth.ExtractUserId(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userId != userIdInToken {
		response.Erro(w, http.StatusUnauthorized, errors.New("user cannot delete another user"))
		return
	}

	//Connect database
	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	if erro = userRepository.DeleteUser(userId); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	followerId, erro := strconv.ParseUint(parameters["followerId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	loggedUserId, erro := auth.ExtractUserId(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	//Connect database
	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	if erro = userRepository.FollowUser(followerId, loggedUserId); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)

	followerId, erro := strconv.ParseUint(parameters["followerId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	loggedUserId, erro := auth.ExtractUserId(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if followerId == loggedUserId {
		response.Erro(w, http.StatusBadRequest, errors.New("you cannot unfollow yourself"))
		return
	}

	//Connect database
	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	if erro = userRepository.UnfollowUser(followerId, loggedUserId); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, nil)
}

func FollowersByUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userId, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	//Connect database
	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	userRepository := repositories.NewUserRepository(db)

	followers, erro := userRepository.FollowersByUser(userId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

func Following(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userId, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	//Connect database
	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	userRepository := repositories.NewUserRepository(db)

	followers, erro := userRepository.Following(userId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userId, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	userIdInToken, erro := auth.ExtractUserId(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userIdInToken != userId {
		response.Erro(w, http.StatusForbidden, errors.New("You cannot change another user password"))
		return
	}

	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	//transform body into a user entity
	var password models.Password
	if erro = json.Unmarshal(bodyRequest, &password); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//Connect database
	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	passwordDB, erro := userRepository.FindPassword(userId)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if erro = security.ComparePassword(passwordDB, password.CurrentPassword); erro != nil {
		response.Erro(w, http.StatusUnauthorized, errors.New("current password is not correct"))
		return
	}

	passwordWithHash, erro := security.Hash(password.NewPassword)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = userRepository.UpdatePassword(userId, string(passwordWithHash)); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, nil)
}
