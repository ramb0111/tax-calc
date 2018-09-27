package main

import (
	"github.com/gorilla/mux"
	"github.com/ramb0111/tax-calc/server"
	"github.com/ramb0111/tax-calc/src/config"
	"github.com/ramb0111/tax-calc/src/routes"
)

func main() {
	r := routes.Route{mux.NewRouter()}
	r.AddHandles()

	server := server.Server{config.Port, config.ErrSleepTime}
	server.Run(&r)
}
