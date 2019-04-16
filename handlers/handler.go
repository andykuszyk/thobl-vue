package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct {
	Router *mux.Router
}

func BuildHandlers() *handler {
	h := &handler{
		Router: mux.NewRouter(),
	}
	h.Router.HandleFunc("/users", h.postUsersHandler).Methods(http.MethodPost)
	return h
}
