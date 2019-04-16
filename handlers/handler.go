package handlers

import (

	"github.com/gorilla/mux"
)

type handler struct {
	Router *mux.Router
}

func BuildHandlers() *handler {
	return &handler{
		Router: mux.NewRouter(),
	}
}
