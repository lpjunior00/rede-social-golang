package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

type User struct {
	Id           uint64    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Nickname     string    `json:"nickname"`
	CreationDate time.Time `json:"creationDate"`
	Followers    []User    `json:"followers"`
	Following    []User    `json:"following"`
	Posts        []Post    `json:"posts"`
}

func LoadUser(userId uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followerChanel := make(chan []User)
	followingChanel := make(chan []User)
	postChanel := make(chan []Post)

	go SearchBasicData(userChannel, userId, r)
	go SearchFollowers(followerChanel, userId, r)
	go SearchFollowing(followingChanel, userId, r)
	go SearchPosts(postChanel, userId, r)

	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	for indice := 0; indice < 4; indice++ {
		//Select é um switch para concorrencia
		select {
		//Se chegou um valor no canal (só pode chegar 1, então faz a ação)
		case loadedUser := <-userChannel:
			if loadedUser.Id == 0 {
				return User{}, errors.New("erro ao buscar os dados básicos do usuário")
			}

			user = loadedUser
		case loadedFollowers := <-followerChanel:
			if loadedFollowers == nil {
				return User{}, errors.New("erro ao buscar os dados dos seguidores do usuário")
			}

			followers = loadedFollowers

		case loadedFollowing := <-followingChanel:
			if loadedFollowing == nil {
				return User{}, errors.New("erro ao buscar os dados dos usuários que está seguindo")
			}

			following = loadedFollowing

		case loadedPost := <-postChanel:
			if loadedPost == nil {
				return User{}, errors.New("erro ao buscar os posts do usuário")
			}

			posts = loadedPost

		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil

}

func SearchBasicData(channel chan<- User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, userId)
	responseApi, erro := requests.RequestWithAutentication(r, http.MethodGet, url, nil)
	if erro != nil {
		channel <- User{}
		return
	}

	defer responseApi.Body.Close()

	var user User
	if erro = json.NewDecoder(responseApi.Body).Decode(&user); erro != nil {
		channel <- User{}
		return
	}

	channel <- user
}

func SearchFollowers(channel chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.ApiUrl, userId)
	responseApi, erro := requests.RequestWithAutentication(r, http.MethodGet, url, nil)
	if erro != nil {
		channel <- nil
		return
	}

	defer responseApi.Body.Close()

	var followers []User
	if erro = json.NewDecoder(responseApi.Body).Decode(&followers); erro != nil {
		channel <- nil
		return
	}

	if followers == nil {
		//mandar um slice vazio porque nao tem seguidores
		channel <- make([]User, 0)
		return
	}

	channel <- followers
}

// Esse parametro apontando pro canal, chan<- []User, quer dizer que vai ser um canal que vai receber informacoes do tipo informado. A ordem da seta determina
func SearchFollowing(channel chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.ApiUrl, userId)
	responseApi, erro := requests.RequestWithAutentication(r, http.MethodGet, url, nil)
	if erro != nil {
		channel <- nil
		return
	}

	defer responseApi.Body.Close()

	var following []User
	if erro = json.NewDecoder(responseApi.Body).Decode(&following); erro != nil {
		channel <- nil
		return
	}

	if following == nil {
		//mandar um slice vazio porque nao está seguindo ninguem
		channel <- make([]User, 0)
		return
	}

	channel <- following
}

func SearchPosts(channel chan<- []Post, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.ApiUrl, userId)
	responseApi, erro := requests.RequestWithAutentication(r, http.MethodGet, url, nil)
	if erro != nil {
		channel <- nil
		return
	}

	defer responseApi.Body.Close()

	var posts []Post
	if erro = json.NewDecoder(responseApi.Body).Decode(&posts); erro != nil {
		channel <- nil
		return
	}

	if posts == nil {
		//mandar um slice vazio porque nao tem posts
		channel <- make([]Post, 0)
		return
	}

	channel <- posts
}
