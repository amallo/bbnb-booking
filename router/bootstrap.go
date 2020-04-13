package router

import (
	"bbnb-booking/auth/handlers"
	repository "bbnb-booking/auth/repository/mongodb"
	"bbnb-booking/auth/session"
	"bbnb-booking/auth/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func Boostrap(secret string) *mux.Router {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	usersRouter := apiRouter.PathPrefix("/users").Subrouter()

	signInUseCase := usecase.SignIn(repository.FindUser, session.CreateWithSecret(secret))
	signInHandler := handlers.SignIn(signInUseCase)
	usersRouter.HandleFunc("/signIn", signInHandler).Methods(http.MethodPost)
	return router
}
