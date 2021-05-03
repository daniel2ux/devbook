package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

//Generate return new route
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
