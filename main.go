package main

import (
	"bbnb-booking/auth/handlers"
	"bbnb-booking/auth/repository"
	"bbnb-booking/auth/session"
	"bbnb-booking/auth/usecase"
	"bbnb-booking/config"
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
)

func main() {
	getConfig := config.WithViperConfig(".env")

	/** Connect to the mongo database **/
	clientOptions := options.Client().ApplyURI(getConfig("MONGO_URL"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("booking")
	fmt.Println("Connected to MongoDB!")

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	usersRouter := apiRouter.PathPrefix("/users").Subrouter()

	findUserByEmailAndPassword := repository.FindUser(database)
	signInUseCase := usecase.SignIn(findUserByEmailAndPassword, session.CreateWithSecret("secret"))
	signInHandler := handlers.SignIn(signInUseCase)
	usersRouter.HandleFunc("/signIn", signInHandler).Methods(http.MethodPost)

	http.ListenAndServe(":80", router)
}
