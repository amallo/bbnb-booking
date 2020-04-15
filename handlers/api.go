package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

func ApiHandler(fn ApiHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json, err := fn(w, r)
		if err != nil {
			http.Error(w, err.Message, err.Code)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}
func DecodePayload(body io.ReadCloser, payload interface{}) *ApiHandlerError {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&payload)
	if err != nil {
		return &ApiHandlerError{Code: http.StatusBadRequest, Message: "Cannot decode payload", Error: err}
	}
	return nil
}

func EncodeResponse(response interface{}) ([]byte, *ApiHandlerError) {
	json, err := json.Marshal(response)
	if err != nil {
		return nil, &ApiHandlerError{Code: http.StatusInternalServerError, Message: "Cannot encode response", Error: err}
	}
	return json, nil
}
