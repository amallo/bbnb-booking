package handlers

import (
	"bbnb-booking/auth/usecase"
	"bbnb-booking/models"
	"encoding/json"
	"net/http"
)

type signInPayload struct {
	Email    string
	Password string
}

type signInResponse struct {
	User          *models.User
	Authorization *string
	Message       string
}

func SignIn(signIn usecase.SignInFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/**
		  Decode payload
		**/
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		var payload signInPayload
		err := decoder.Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		/**
			Signin with email and password
		**/
		token, user, err := signIn(usecase.Credentials{Email: payload.Email, Password: payload.Password})
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		response := signInResponse{User: user, Authorization: token, Message: "Ok"}
		json, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}
