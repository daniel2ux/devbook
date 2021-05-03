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
		IsRequiredAuth: false,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodGet,
		Function:       controllers.GetUser,
		IsRequiredAuth: false,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodPut,
		Function:       controllers.UpdateUser,
		IsRequiredAuth: false,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodDelete,
		Function:       controllers.DeleteUser,
		IsRequiredAuth: false,
	},
}
