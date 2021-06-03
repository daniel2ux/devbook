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
}
