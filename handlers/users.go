package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/andykuszyk/thobl/models"
)

func (h *handler) postUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func (h *handler) postUsersAuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	requestUser := models.User{}
	err := json.NewDecoder(r.Body).Decode(&requestUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := h.UsersRepo.GetByUsername(requestUser.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if user.Password != requestUser.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	userToken := models.UserToken {
		Token: "foo",
	}
	bytes, err := json.Marshal(userToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
