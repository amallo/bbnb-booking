package usecase

import "bbnb-booking/models"

type Credentials struct {
	Email    string
	Password string
}

type SignInFunc = func(Credentials) (*string, *models.User, error)
