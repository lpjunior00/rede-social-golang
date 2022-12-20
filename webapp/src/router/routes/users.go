package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var userRoutes = []Route{
	{
		Uri:            "/create-user",
		Method:         http.MethodGet,
		Function:       controllers.LoadCreateUserPage,
		isAutenticated: false,
	},
	{
		Uri:            "/users",
		Method:         http.MethodPost,
		Function:       controllers.CreateUser,
		isAutenticated: false,
	},
	{
		Uri:            "/users/search",
		Method:         http.MethodGet,
		Function:       controllers.LoadUserPage,
		isAutenticated: false,
	},
	{
		Uri:            "/users/{userId}",
		Method:         http.MethodGet,
		Function:       controllers.LoadUserDetails,
		isAutenticated: false,
	},
	{
		Uri:            "/users/{userId}/unfollow",
		Method:         http.MethodPost,
		Function:       controllers.UnfollowUser,
		isAutenticated: false,
	},
	{
		Uri:            "/users/{userId}/follow",
		Method:         http.MethodPost,
		Function:       controllers.FollowUser,
		isAutenticated: false,
	},
	{
		Uri:            "/profile",
		Method:         http.MethodGet,
		Function:       controllers.LoadProfilePage,
		isAutenticated: false,
	},
	{
		Uri:            "/edit-user",
		Method:         http.MethodGet,
		Function:       controllers.LoadEditUserPage,
		isAutenticated: false,
	},
	{
		Uri:            "/edit-user",
		Method:         http.MethodPut,
		Function:       controllers.EditUser,
		isAutenticated: false,
	},
	{
		Uri:            "/change-password",
		Method:         http.MethodGet,
		Function:       controllers.LoadChangePasswordPage,
		isAutenticated: false,
	},
	{
		Uri:            "/change-password",
		Method:         http.MethodPost,
		Function:       controllers.ChangePassword,
		isAutenticated: false,
	},
	{
		Uri:            "/users",
		Method:         http.MethodDelete,
		Function:       controllers.DeleteUser,
		isAutenticated: false,
	},
}
