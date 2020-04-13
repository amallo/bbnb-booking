package api

import (
	"bbnb-booking/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	routes.HandleSignIn(router, "secret")
	router.ServeHTTP(w, r)
}
