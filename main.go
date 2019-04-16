package main

import (
	"net/http"

	"github.com/andykuszyk/thobl/handlers"
)

func main() {
	server := &http.Server{
		Handler: handlers.BuildHandlers().Router,
		Addr: ":8080",
	}
	server.ListenAndServe();
}
