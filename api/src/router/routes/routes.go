package routes

import (
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

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
