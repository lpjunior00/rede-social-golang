package controllers

import (
	"net/http"
)

func SetupDatabase(w http.ResponseWriter, r *http.Request) {
	// db, erro := database.Connect()
	// if erro != nil {
	// 	response.Erro(w, http.StatusInternalServerError, erro)
	// 	return
	// }
	// defer db.Close()

	// postRepository := repositories.NewPostRepository(db)
	// posts, erro := postRepository.FindAll(userId)

	// if erro != nil {
	// 	response.Erro(w, http.StatusInternalServerError, erro)
	// 	return
	// }

	// if len(posts) == 0 {
	// 	response.JSON(w, http.StatusOK, []models.Post{})
	// 	return
	// }

	// response.JSON(w, http.StatusOK, posts)
}
