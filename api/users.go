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

	"go.mongodb.org/mongo-driver/mongo"
)

var database *mongo.Database

func SignInHandler(w http.ResponseWriter, r *http.Request) {

	/** Connect to the mongo database **/
	getConfig := config.WithEnvConfig(os.Getenv)
	mongoUri := getConfig("MONGO_URL")
	database, err := data.MemoConnectToDatase(mongoUri, "bookings")

	/** If invalid database handler then stop execution**/
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	findUserByEmailAndPassword := repository.FindUser(database)
	signInUseCase := usecase.SignIn(findUserByEmailAndPassword, session.CreateWithSecret("secret"))
	signInHandler := handlers.SignIn(signInUseCase)
	signInHandler(w, r)
}
