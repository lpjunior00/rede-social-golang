package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return routes.Configure(r)
}
