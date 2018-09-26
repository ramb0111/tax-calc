package routes

import (
	"github.com/gorilla/mux"
	"github.com/ramb0111/tax-calc/src/controllers"
)

type Route struct {
	R *mux.Router
}

func (route *Route) AddHandles() {
	route.R.HandleFunc("/tax/{userId}", controllers.CreateItem).Methods("POST")
	route.R.HandleFunc("/tax/{userId}", controllers.GetItems).Methods("GET")
}
