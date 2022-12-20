package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	isAutenticated bool
}

// Configura todas as rotas dentro do usuario
func Configure(r *mux.Router) *mux.Router {
	routes := routesUsuarios
	routes = append(routes, routeLogin)
	//o Spread no go é feito depois da variavel
	routes = append(routes, routesPosts...)

	for _, route := range routes {

		if route.isAutenticated {
			r.HandleFunc(route.Uri,
				middlewares.Logger(
					middlewares.Authenticate(route.Function))).Methods(route.Method)

		} else {
			r.HandleFunc(route.Uri, middlewares.Logger(route.Function)).Methods(route.Method)

		}

	}

	return r
}
