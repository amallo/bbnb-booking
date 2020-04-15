package main

import (
	"bbnb-booking/api"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	usersRouter := apiRouter.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("/signIn", api.SignInHandler).Methods(http.MethodPost)
	usersRouter.HandleFunc("/signUp", api.SignUpHandler).Methods(http.MethodPost)
	http.ListenAndServe(":80", router)
}
