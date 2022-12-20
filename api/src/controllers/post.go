package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FindPosts(w http.ResponseWriter, r *http.Request) {
	userId, erro := auth.ExtractUserId(r)

	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostRepository(db)
	posts, erro := postRepository.FindAll(userId)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if len(posts) == 0 {
		response.JSON(w, http.StatusOK, []models.Post{})
		return
	}

	response.JSON(w, http.StatusOK, posts)

}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	//transform body into a user entity
	var post models.Post
	if erro = json.Unmarshal(bodyRequest, &post); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	authorId, erro := auth.ExtractUserId(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	post.AuthorId = authorId

	//Validate and fix data from user entity
	if erro := post.Prepare(); erro != nil {
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

	//Create a post repository and save post entity
	postRepository := repositories.NewPostRepository(db)
	postId, erro := postRepository.Create(post)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//return id created
	post.Id = postId
	response.JSON(w, http.StatusCreated, post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	authorId, erro := auth.ExtractUserId(r)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	postRepository := repositories.NewPostRepository(db)
	postDb, erro := postRepository.FindSpecificPost(postId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if postDb.AuthorId != authorId {
		response.Erro(w, http.StatusInternalServerError, errors.New("You cannot delete another user's post"))
		return
	}
	if erro := postRepository.DeletePost(postId); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, nil)

}

func FindSpecificPost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostRepository(db)
	post, erro := postRepository.FindSpecificPost(postId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var post models.Post
	if erro := json.Unmarshal(bodyRequest, &post); erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	authorId, erro := auth.ExtractUserId(r)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	post.AuthorId = authorId

	if erro := post.Prepare(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostRepository(db)
	postDb, erro := postRepository.FindSpecificPost(postId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if postDb.AuthorId != authorId {
		response.Erro(w, http.StatusInternalServerError, errors.New("You cannot update another user's post"))
		return
	}

	if erro := postRepository.UpdatePost(postId, post); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, nil)

}

func FindPostsByUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostRepository(db)
	posts, erro := postRepository.FindPostsByUser(userId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, posts)

}

func LikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	postRepository := repositories.NewPostRepository(db)

	if erro := postRepository.LikePost(postId); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, nil)
}

func UnlikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	postRepository := repositories.NewPostRepository(db)
	if erro := postRepository.UnlikePost(postId); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, nil)
}
