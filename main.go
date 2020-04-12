package main

import (
	httpDeliver "housings-api/auth/delivery/http"
	repository "housings-api/auth/repository/mongodb"
	session "housings-api/auth/session"
	"housings-api/auth/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()
	usersRouter := apiRouter.PathPrefix("/users").Subrouter()
	signInUseCase := usecase.SignIn(repository.FindUser, session.CreateWithSecret("secret"))

	httpDeliver.HandleSignIn(usersRouter, signIn)

	http.ListenAndServe(":80", r)
}
