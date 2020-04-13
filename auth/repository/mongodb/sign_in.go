package repository

import (
	"bbnb-booking/auth/repository"
	"fmt"
)

func FindUser(email string, password string) (*repository.User, error) {
	if email != "audie@app2b.io" || password != "12345" {
		return nil, fmt.Errorf("User %s not matching email or password", email)
	}
	return &repository.User{Email: email}, nil

}
