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
}
