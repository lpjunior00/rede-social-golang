package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesPosts = []Route{
	{
		Uri:            "/posts",
		Method:         http.MethodGet,
		Function:       controllers.FindPosts,
		isAutenticated: false,
	},
	{
		Uri:            "/posts",
		Method:         http.MethodPost,
		Function:       controllers.CreatePost,
		isAutenticated: true,
	},
	{
		Uri:            "/posts/{postId}",
		Method:         http.MethodDelete,
		Function:       controllers.DeletePost,
		isAutenticated: true,
	},
	{
		Uri:            "/posts/{postId}",
		Method:         http.MethodGet,
		Function:       controllers.FindSpecificPost,
		isAutenticated: false,
	},
	{
		Uri:            "/posts/{postId}",
		Method:         http.MethodPut,
		Function:       controllers.UpdatePost,
		isAutenticated: true,
	},
	{
		Uri:            "/users/{userId}/posts",
		Method:         http.MethodGet,
		Function:       controllers.FindPostsByUser,
		isAutenticated: false,
	},
	{
		Uri:            "/posts/{postId}/like",
		Method:         http.MethodPost,
		Function:       controllers.LikePost,
		isAutenticated: true,
	},
	{
		Uri:            "/posts/{postId}/dislike",
		Method:         http.MethodPost,
		Function:       controllers.UnlikePost,
		isAutenticated: true,
	},
}
