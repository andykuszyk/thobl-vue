package handlers

import (
	"net/http"
)

func (h *handler) postUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func (h *handler) postUsersAuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
