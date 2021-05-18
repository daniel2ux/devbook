package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoute = []Route{
	{
		URI:            "/post",
		Method:         http.MethodPost,
		Function:       controllers.NewPost,
		IsRequiredAuth: true,
	},
	{
		URI:            "/post",
		Method:         http.MethodGet,
		Function:       controllers.GetPosts,
		IsRequiredAuth: true,
	},
	{
		URI:            "/post/{id}",
		Method:         http.MethodGet,
		Function:       controllers.GetPost,
		IsRequiredAuth: true,
	},
	{
		URI:            "/post/{id}",
		Method:         http.MethodPut,
		Function:       controllers.UpdatePost,
		IsRequiredAuth: true,
	},
	{
		URI:            "/post/{id}",
		Method:         http.MethodDelete,
		Function:       controllers.DeletePost,
		IsRequiredAuth: true,
	},
}
