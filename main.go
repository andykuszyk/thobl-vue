package main

import (
	"net/http"

	"github.com/andykuszyk/thobl/handlers"
	"github.com/andykuszyk/thobl/repos"
)

func main() {
	usersRepo := &repos.UsersPostgresRepo{}
	server := &http.Server{
		Handler: handlers.BuildHandlers(usersRepo).Router,
		Addr: ":8080",
	}
	server.ListenAndServe();
}
