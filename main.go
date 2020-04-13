package main

import (
	"bbnb-booking/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	usersRouter := apiRouter.PathPrefix("/users").Subrouter()
	routes.HandleSignIn(usersRouter, "/signIn", "secret")
	http.ListenAndServe(":80", router)
}
