package router

import (
	"github.com/kerimkuscu/hardline-fitness-backend/api/router/routes"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetUpRoutesWithMiddlewares(r)
}
