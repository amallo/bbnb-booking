package handlers

import "net/http"

type ApiHandlerError struct {
	Error   error
	Message string
	Code    int
}
type ApiHandlerFunc func(http.ResponseWriter, *http.Request) ([]byte, *ApiHandlerError)
