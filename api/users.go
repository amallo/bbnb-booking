package api

import (
	"bbnb-booking/auth/handlers"
	repository "bbnb-booking/auth/repository/mongodb"
	"bbnb-booking/auth/session"
	"bbnb-booking/auth/usecase"
	"net/http"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	signInUseCase := usecase.SignIn(repository.FindUser, session.CreateWithSecret("secret"))
	signInHandler := handlers.SignIn(signInUseCase)
	signInHandler(w, r)
}
