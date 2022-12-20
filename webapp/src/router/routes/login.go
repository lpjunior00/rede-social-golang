package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routesLogin = []Route{
	{
		Uri:            "/",
		Method:         http.MethodGet,
		Function:       controllers.LoadLoginPage,
		isAutenticated: false,
	},
	{
		Uri:            "/login",
		Method:         http.MethodGet,
		Function:       controllers.LoadLoginPage,
		isAutenticated: false,
	},
	{
		Uri:            "/login",
		Method:         http.MethodPost,
		Function:       controllers.Login,
		isAutenticated: false,
	},
}
