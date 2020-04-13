package main

import (
	"bbnb-booking/router"
	"net/http"
)

func main() {
	router := router.Boostrap("secret")
	http.ListenAndServe(":80", router)
}
