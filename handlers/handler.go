package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/andykuszyk/thobl/repos"
)

type handler struct {
	Router *mux.Router
	UsersRepo repos.UsersRepo
}

func BuildHandlers(usersRepo repos.UsersRepo) *handler {
	h := &handler{
		Router: mux.NewRouter(),
		UsersRepo: usersRepo,
	}
	h.Router.HandleFunc("/users", h.postUsersHandler).Methods(http.MethodPost)
	h.Router.HandleFunc("/users/authenticate", h.postUsersAuthenticateHandler).Methods(http.MethodPost)
	return h
}
