package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/pkg/errors"
	"github.com/ramb0111/tax-calc/src/routes"
)

type Server struct {
	Port         string
	ErrSleepTime int
}

func (s *Server) Run(route *routes.Route) {
	log.Println("[API Server] Port:", s.Port)
	for err := http.ListenAndServe(":"+s.Port, handlers.LoggingHandler(os.Stdout, route.R)); err != nil; {

		log.Printf("[API Server] Err: %+v", errors.Wrap(err, ""))
		time.Sleep(time.Duration(s.ErrSleepTime) * time.Millisecond)

		log.Println("[API Server] Retry Port:", s.Port)
		err = http.ListenAndServe(":"+s.Port, handlers.LoggingHandler(os.Stdout, route.R))
	}
}
