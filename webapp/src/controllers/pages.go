package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/response"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookie(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
	} else {
		utils.ExecuteTemplate(w, "login.html", nil)
	}
}

func LoadCreateUserPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "createUser.html", nil)
}

func LoadHomePage(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("%s/posts", config.ApiUrl)
	responseReq, erro := requests.RequestWithAutentication(r, http.MethodGet, url, nil)

	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

	if responseReq.StatusCode >= 400 {
		response.ErrorHandler(w, responseReq)
		return
	}

	defer responseReq.Body.Close()

	var posts []models.Post
	if erro = json.NewDecoder(responseReq.Body).Decode(&posts); erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.ReadCookie(r)
	userId, _ := strconv.ParseUint(cookie["userId"], 10, 64)

	utils.ExecuteTemplate(w, "home.html", struct {
		Posts  []models.Post
		UserId uint64
	}{
		Posts:  posts,
		UserId: userId,
	})
}

func LoadEditPostPage(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.ApiUrl, postId)
	responseReq, erro := requests.RequestWithAutentication(r, http.MethodGet, url, nil)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

	if responseReq.StatusCode >= 400 {
		response.ErrorHandler(w, responseReq)
		return
	}

	defer responseReq.Body.Close()

	var post models.Post
	if erro = json.NewDecoder(responseReq.Body).Decode(&post); erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecuteTemplate(w, "editPost.html", post)

}

func LoadUserPage(w http.ResponseWriter, r *http.Request) {
	nickname := strings.ToLower(r.URL.Query().Get("userNickname"))

	url := fmt.Sprintf("%s/users?user=%s", config.ApiUrl, nickname)
	fmt.Println(url)

	responseReq, erro := requests.RequestWithAutentication(r, http.MethodGet, url, nil)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

	if responseReq.StatusCode >= 400 {
		response.ErrorHandler(w, responseReq)
		return
	}

	var users []models.User
	if erro := json.NewDecoder(responseReq.Body).Decode(&users); erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecuteTemplate(w, "users.html", users)

}

func LoadUserDetails(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.ReadCookie(r)
	loggedUserId, _ := strconv.ParseUint(cookie["userId"], 10, 64)

	if userId == loggedUserId {
		http.Redirect(w, r, "/profile", http.StatusFound)
	}

	user, erro := models.LoadUser(userId, r)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecuteTemplate(w, "user.html", struct {
		User         models.User
		LoggedUserId uint64
	}{
		User:         user,
		LoggedUserId: loggedUserId,
	})

}

func LoadProfilePage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookie(r)
	loggedUserId, _ := strconv.ParseUint(cookie["userId"], 10, 64)

	user, erro := models.LoadUser(loggedUserId, r)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecuteTemplate(w, "profile.html", user)
}

func LoadEditUserPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookie(r)
	loggedUserId, _ := strconv.ParseUint(cookie["userId"], 10, 64)

	channel := make(chan models.User)
	go models.SearchBasicData(channel, loggedUserId, r)
	user := <-channel

	if user.Id == 0 {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: "Erro ao buscar o usuário"})
		return
	}

	utils.ExecuteTemplate(w, "editUser.html", user)
}

func LoadChangePasswordPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "changePassword.html", nil)
}
