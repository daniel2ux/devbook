package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

//Route represents all routes of the application
type Route struct {
	URI            string
	Method         string
	Function       func(w http.ResponseWriter, r *http.Request)
	IsRequiredAuth bool
}

//Config return all configureds routes
func Config(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, postsRoute...)

	for _, route := range routes {
		if route.IsRequiredAuth {
			r.HandleFunc(
				route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
