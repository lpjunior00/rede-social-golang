package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Escreve algmas informacoes da chamada no terminal
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)

		nextFunc(w, r)
	}
}

// Verifica se o cookie existe (a validação é na API)
func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, erro := cookies.ReadCookie(r); erro != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		nextFunc(w, r)
	}
}
