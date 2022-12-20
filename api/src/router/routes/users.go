package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsuarios = []Route{
	{
		Uri:            "/users",
		Method:         http.MethodPost,
		Function:       controllers.CreateUser,
		isAutenticated: false,
	},
	{
		Uri:            "/users",
		Method:         http.MethodGet,
		Function:       controllers.FindUsers,
		isAutenticated: true,
	},
	{
		Uri:            "/users/{userId}",
		Method:         http.MethodGet,
		Function:       controllers.FindSpecificUser,
		isAutenticated: false,
	},
	{
		Uri:            "/users/{userId}",
		Method:         http.MethodPut,
		Function:       controllers.UpdateUser,
		isAutenticated: true,
	},
	{
		Uri:            "/users/{userId}",
		Method:         http.MethodDelete,
		Function:       controllers.DeleteUser,
		isAutenticated: true,
	},
	{
		Uri:            "/users/follow/{followerId}",
		Method:         http.MethodPost,
		Function:       controllers.FollowUser,
		isAutenticated: true,
	},
	{
		Uri:            "/users/unfollow/{followerId}",
		Method:         http.MethodPost,
		Function:       controllers.UnfollowUser,
		isAutenticated: true,
	},
	{
		Uri:            "/users/{userId}/followers",
		Method:         http.MethodGet,
		Function:       controllers.FollowersByUser,
		isAutenticated: false,
	},
	{
		Uri:            "/users/{userId}/following",
		Method:         http.MethodGet,
		Function:       controllers.Following,
		isAutenticated: false,
	},
	{
		Uri:            "/users/{userId}/update-password",
		Method:         http.MethodPost,
		Function:       controllers.UpdatePassword,
		isAutenticated: true,
	},
}
