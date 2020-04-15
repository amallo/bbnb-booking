package api

import (
	authHandlers "bbnb-booking/auth/handlers"
	"bbnb-booking/auth/repository"
	"bbnb-booking/auth/session"
	"bbnb-booking/auth/usecase"
	"bbnb-booking/config"
	"bbnb-booking/data"
	handlers "bbnb-booking/handlers"
	"net/http"
)

func signInHandler(w http.ResponseWriter, r *http.Request) ([]byte, *handlers.ApiHandlerError) {

	/** Connect to the mongo database **/
	conf := config.GetEnvConfig()
	database, err := data.MemoConnectToDatase(conf.DatabaseURI, "bookings")
	if err != nil {
		return nil, &handlers.ApiHandlerError{Code: http.StatusInternalServerError, Error: err, Message: "Unable to connect to database"}
	}

	findUser := repository.FindUser(database)
	createSession := session.CreateWithSecret("secret")
	signInUseCase := usecase.SignInUseCase(findUser, createSession)
	signInHandler := authHandlers.SignInHandler(signInUseCase)
	return signInHandler(w, r)
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	handlers.ApiHandler(signInHandler)(w, r)
}
