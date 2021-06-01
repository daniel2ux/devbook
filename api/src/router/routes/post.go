package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoute = []Route{
	{
		URI:            "/posts",
		Method:         http.MethodPost,
		Function:       controllers.NewPost,
		IsRequiredAuth: true,
	},
	{
		URI:            "/posts",
		Method:         http.MethodGet,
		Function:       controllers.GetPosts,
		IsRequiredAuth: true,
	},
	{
		URI:            "/posts/{id}",
		Method:         http.MethodGet,
		Function:       controllers.GetPost,
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
	{
		URI:            "/users/{id}/posts",
		Method:         http.MethodGet,
		Function:       controllers.GetPostsByUser,
		IsRequiredAuth: true,
	},
}
