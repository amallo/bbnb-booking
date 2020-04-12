package main

import (
	deliver "housings-api/auth/delivery/http"
	"housings-api/auth/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()
	usersRouter := apiRouter.PathPrefix("/users").Subrouter()

	deliver.HandleSignIn(usersRouter, usecase.SignIn)

	http.ListenAndServe(":80", r)
}
