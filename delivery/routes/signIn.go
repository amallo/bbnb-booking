package delivery

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type signInPayload struct {
	Email    string
	Password string
}

type signInResponse struct {
	Authorization string
	Message       string
}

func SignIn(w http.ResponseWriter, r *http.Request) {
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
		Generate a new token from email
	**/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":   "housings-users-api",                                  // issuer
		"sub":   "auth",                                                // subject
		"aud":   "webstart students",                                   // audience
		"email": payload.Email,                                         // extra info
		"nbf":   time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(), // not before
	})
	// md5 password ?
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	loginResponse := signInResponse{tokenString, "OK"}
	json, err := json.Marshal(loginResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
