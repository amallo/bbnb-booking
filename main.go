package main

import (
	routes "housings-api/delivery/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()
	usersRouter := apiRouter.PathPrefix("/users").Subrouter()

	usersRouter.HandleFunc("/signIn", routes.SignIn).Methods(http.MethodPost)
	http.ListenAndServe(":80", r)
}
