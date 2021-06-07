package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	IsRequiredAuth bool
}

func Config(router *mux.Router) *mux.Router {
	routes := routesLogin
	routes = append(routes, userRoutes...)
	routes = append(routes, mainPageRoute)

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
	return router
}
