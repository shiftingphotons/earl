package app

import (
	"earl/app/handlers"
	"earl/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func (server *Server) Initialize() {
	db.Connect()

	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/api/ping", handlers.Ping).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/shorten", handlers.Shorten).Methods(http.MethodPost)
	server.Router.HandleFunc("/{shortUrl:[a-zA-Z0-9]+}", handlers.Redirect).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8000", server.Router))
}
