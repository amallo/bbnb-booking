package api

import (
	repository "housings-api/auth/repository/mongodb"
	"housings-api/auth/session"
	"housings-api/auth/usecase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting api...")
	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()
	usersRouter := apiRouter.PathPrefix("/users").Subrouter()
	signInUseCase := usecase.SignIn(repository.FindUser, session.CreateWithSecret("secret"))

	httpDeliver.HandleSignIn(usersRouter, signInUseCase)

	http.ListenAndServe(":80", r)
}
