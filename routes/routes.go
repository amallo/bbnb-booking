package routes

import (
	"bbnb-booking/auth/handlers"
	repository "bbnb-booking/auth/repository/mongodb"
	"bbnb-booking/auth/session"
	"bbnb-booking/auth/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleSignIn(router *mux.Router, routeName string, secret string) {
	signInUseCase := usecase.SignIn(repository.FindUser, session.CreateWithSecret(secret))
	signInHandler := handlers.SignIn(signInUseCase)
	router.HandleFunc(routeName, signInHandler).Methods(http.MethodPost)
}
