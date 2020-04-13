package handlers

import (
	"net/http"
)

type HandlerFunc = func(w http.ResponseWriter, r *http.Request)
