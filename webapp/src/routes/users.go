package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var userRoutes = []Route{
	{
		URI:            "/enroll",
		Method:         http.MethodGet,
		Function:       controllers.LoadUserEnrollPage,
		IsRequiredAuth: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Function:       controllers.EnrollUser,
		IsRequiredAuth: false,
	},
	{
		URI:            "/get-users",
		Method:         http.MethodGet,
		Function:       controllers.LoadUsersPage,
		IsRequiredAuth: true,
	},
}
