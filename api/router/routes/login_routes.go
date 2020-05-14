package routes

import (
	"github.com/kerimkuscu/hardline-fitness-backend/api/controllers"
	"net/http"
)

var loginRoutes = []Route{
	Route{
		Uri:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
}
