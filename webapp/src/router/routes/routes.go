package routes

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	isAutenticated bool
}

// coloca todas as rotas para dentro do router
func Configure(router *mux.Router) *mux.Router {
	routes := routesLogin
	routes = append(routes, userRoutes...)
	routes = append(routes, homePageRoute)
	routes = append(routes, postRoutes...)
	routes = append(routes, logoutRoutes)

	for _, route := range routes {

		if route.isAutenticated {
			router.HandleFunc(route.Uri, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			router.HandleFunc(route.Uri, middlewares.Logger(route.Function)).Methods(route.Method)
		}

	}

	//Define a pasta que contem os assets (css e js). Sem isso o GO não consegue carregar os estilos
	fileServer := http.FileServer(http.Dir("./assets/"))
	//Aqui define a pasta raiz dos assets ,pra não fica tendo que ficar colocando pontos pra descer de diretorio no frontend
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fileServer))

	return router
}
