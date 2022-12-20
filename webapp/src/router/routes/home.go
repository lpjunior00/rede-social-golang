package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var homePageRoute = Route{
	Uri:            "/home",
	Method:         http.MethodGet,
	Function:       controllers.LoadHomePage,
	isAutenticated: true,
}
