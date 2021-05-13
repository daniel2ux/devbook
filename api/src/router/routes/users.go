package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Function:       controllers.CreateUser,
		IsRequiredAuth: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Function:       controllers.GetUsers,
		IsRequiredAuth: true,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodGet,
		Function:       controllers.GetUser,
		IsRequiredAuth: true,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodPut,
		Function:       controllers.UpdateUser,
		IsRequiredAuth: true,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodDelete,
		Function:       controllers.DeleteUser,
		IsRequiredAuth: true,
	},
	{
		URI:            "/users/{id}/follow",
		Method:         http.MethodPost,
		Function:       controllers.FollowUser,
		IsRequiredAuth: true,
	},
	{
		URI:            "/users/{id}/stop-follow",
		Method:         http.MethodPost,
		Function:       controllers.StopFollowUser,
		IsRequiredAuth: true,
	},
	{
		URI:            "/users/{id}/followers",
		Method:         http.MethodGet,
		Function:       controllers.GetFollowers,
		IsRequiredAuth: true,
	},
}
