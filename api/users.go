package api

import (
	"bbnb-booking/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	apiRouter := mux.NewRouter()
	routes.HandleSignIn(apiRouter, "/signIn", "secret")
	apiRouter.ServeHTTP(w, r)
}
