package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(router *mux.Router) {
	StudentRoutes(router)
}