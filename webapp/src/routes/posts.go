package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var postRoutes = []Route{
	{
		URI:            "/posts",
		Method:         http.MethodPost,
		Function:       controllers.NewPost,
		IsRequiredAuth: true,
	},
	{
		URI:            "/posts/{id}/like",
		Method:         http.MethodPost,
		Function:       controllers.LikePost,
		IsRequiredAuth: true,
	},
	{
		URI:            "/posts/{id}/dislike",
		Method:         http.MethodPost,
		Function:       controllers.DislikePost,
		IsRequiredAuth: true,
	},
	{
		URI:            "/posts/{id}/edit",
		Method:         http.MethodGet,
		Function:       controllers.LoadEditPage,
		IsRequiredAuth: true,
	},
	{
		URI:            "/posts/{id}",
		Method:         http.MethodPut,
		Function:       controllers.UpdatePost,
		IsRequiredAuth: true,
	},
	{
		URI:            "/posts/{id}",
		Method:         http.MethodDelete,
		Function:       controllers.DeletePost,
		IsRequiredAuth: true,
	},
}
