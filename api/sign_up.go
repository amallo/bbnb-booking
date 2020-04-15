package api

import (
	"bbnb-booking/auth/handlers"
	"bbnb-booking/auth/repository"
	"bbnb-booking/auth/session"
	"bbnb-booking/auth/usecase"
	"bbnb-booking/config"
	"bbnb-booking/data"
	"log"
	"net/http"
	"os"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	/** Connect to the mongo database **/
	getConfig := config.WithEnvConfig(os.Getenv)
	mongoUri := getConfig("MONGO_URL")
	database, err := data.MemoConnectToDatase(mongoUri, "bookings")
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	findUserByEmailAndPassword := repository.FindUser(database)
	insertUser := repository.InsertUser(database)
	useCase := usecase.SignUpUseCase(findUserByEmailAndPassword, insertUser, session.CreateWithSecret("secret"))
	handler := handlers.SignUpHandler(useCase)
	handler(w, r)
}
