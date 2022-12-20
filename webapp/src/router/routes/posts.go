package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var postRoutes = []Route{
	{
		Uri:            "/posts",
		Method:         http.MethodPost,
		Function:       controllers.CreatePost,
		isAutenticated: true,
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
		Function:       controllers.DislikePost,
		isAutenticated: true,
	},
	{
		Uri:            "/posts/{postId}/edit",
		Method:         http.MethodGet,
		Function:       controllers.LoadEditPostPage,
		isAutenticated: true,
	},
	{
		Uri:            "/posts/{postId}",
		Method:         http.MethodPut,
		Function:       controllers.UpdatePost,
		isAutenticated: true,
	},
	{
		Uri:            "/posts/{postId}",
		Method:         http.MethodDelete,
		Function:       controllers.DeletePost,
		isAutenticated: true,
	},
}
