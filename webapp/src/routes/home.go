package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var mainPageRoute = Route{
	URI:            "/home",
	Method:         http.MethodGet,
	Function:       controllers.LoadMainPage,
	IsRequiredAuth: true,
}
