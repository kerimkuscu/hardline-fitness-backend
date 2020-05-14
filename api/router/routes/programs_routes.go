package routes

import (
	"github.com/kerimkuscu/hardline-fitness-backend/api/controllers"
	"net/http"
)

var programsRoutes = []Route{
	Route{
		Uri:          "/programs",
		Method:       http.MethodGet,
		Handler:      controllers.GetPrograms,
		AuthRequired: false,
	},
	Route{
		Uri:          "/programs",
		Method:       http.MethodPost,
		Handler:      controllers.CreateProgram,
		AuthRequired: true,
	},
	Route{
		Uri:          "/programs/{id}",
		Method:       http.MethodGet,
		Handler:      controllers.GetProgram,
		AuthRequired: false,
	},
	Route{
		Uri:          "/programs/{id}",
		Method:       http.MethodPut,
		Handler:      controllers.UpdateProgram,
		AuthRequired: true,
	},
	Route{
		Uri:          "/programs/{id}",
		Method:       http.MethodDelete,
		Handler:      controllers.DeleteProgram,
		AuthRequired: true,
	},
}
