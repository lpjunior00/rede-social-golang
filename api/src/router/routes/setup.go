package routes

import (
	"api/src/controllers"
	"net/http"
)

var setupRoutes = Route{
	Uri:            "/setup",
	Method:         http.MethodGet,
	Function:       controllers.SetupDatabase,
	isAutenticated: false,
}
