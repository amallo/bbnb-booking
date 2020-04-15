package handlers

import (
	"bbnb-booking/auth/usecase"
	"bbnb-booking/handlers"
	"bbnb-booking/models"
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

func SignInHandler(signIn usecase.AuthFunc) handlers.ApiHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) ([]byte, *handlers.ApiHandlerError) {
		/**
		  Decode payload
		**/
		payload := userPayload{}
		errDecode := handlers.DecodePayload(r.Body, &payload)
		if errDecode != nil {
			return nil, errDecode
		}

		/**
			Signin with email and password
		**/
		token, user, err := signIn(usecase.Credentials{Email: payload.Email, Password: payload.Password})
		if err != nil {
			return nil, &handlers.ApiHandlerError{Code: http.StatusForbidden, Message: err.Error(), Error: err}
		}

		/**
			Return response
		**/
		response := userResponse{User: user, Authorization: token}
		json, errEncode := handlers.EncodeResponse(response)
		if errEncode != nil {
			return nil, errEncode
		}
		return json, nil
	}
}

func SignUpHandler(signup usecase.AuthFunc) handlers.ApiHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) ([]byte, *handlers.ApiHandlerError) {
		/**
		  Decode payload
		**/
		payload := userPayload{}
		errDecode := handlers.DecodePayload(r.Body, &payload)
		if errDecode != nil {
			return nil, errDecode
		}

		/**
			Signup with email and password
		**/
		token, user, errUseCase := signup(usecase.Credentials{Email: payload.Email, Password: payload.Password})
		if errUseCase != nil {
			return nil, &handlers.ApiHandlerError{Code: http.StatusForbidden, Message: errUseCase.Error(), Error: errUseCase}
		}

		/**
			Return response
		**/
		response := userResponse{User: user, Authorization: token}
		json, errEncode := handlers.EncodeResponse(response)
		if errEncode != nil {
			return nil, errEncode
		}
		return json, nil
	}
}
