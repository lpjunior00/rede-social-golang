package middlewares

import (
	"api/src/auth"
	"api/src/response"
	"fmt"
	"log"
	"net/http"
)

// Aqui ele recebe uma function, faz uma acao e depois chama a func. é uma especie de proxy
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(" \n %s - %s - %s", r.Method, r.RequestURI, r.Host)

		nextFunction(w, r)
	}

}

// Autenticar verifica se o usuario está autenticado no sistema. É um midleware, que fica entre as requests
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Autenticando")

		if erro := auth.ValidateToken(r); erro != nil {
			response.Erro(w, http.StatusUnauthorized, erro)
			return
		}

		nextFunction(w, r)
	}
}
