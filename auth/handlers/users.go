package handlers

import (
	"bbnb-booking/auth/usecase"
	"bbnb-booking/models"
	"encoding/json"
	"net/http"
)

type userPayload struct {
	Email    string
	Password string
}

type userResponse struct {
	User          *models.User `json:"user"`
	Authorization *string      `json:"authorization"`
}

func SignInHandler(signIn usecase.AuthFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/**
		  Decode payload
		**/
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		payload := userPayload{}
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

		response := userResponse{User: user, Authorization: token}
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

func SignUpHandler(signup usecase.AuthFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/**
		  Decode payload
		**/
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		payload := userPayload{}
		err := decoder.Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		/**
			Signup with email and password
		**/
		token, user, err := signup(usecase.Credentials{Email: payload.Email, Password: payload.Password})
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		response := userResponse{User: user, Authorization: token}
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
