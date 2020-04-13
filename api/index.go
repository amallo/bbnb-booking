package api

import (
	"bbnb-booking/router"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := router.Boostrap("secret")
	router.ServeHTTP(w, r)
}
