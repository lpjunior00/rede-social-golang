package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var logoutRoutes = Route{
	Uri:            "/logout",
	Method:         http.MethodGet,
	Function:       controllers.Logout,
	isAutenticated: true,
}
