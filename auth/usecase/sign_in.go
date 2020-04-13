package usecase

import (
	repository "bbnb-booking/auth/repository"
	session "bbnb-booking/auth/session"
)

func SignIn(findUser repository.FindUserFunc, signSession session.SignSessionFunc) SignInFunc {
	return func(credentials Credentials) (*string, error) {
		user, err := findUser(credentials.Email, credentials.Password)
		if err != nil {
			return nil, err
		}
		return signSession(user.Email)
	}
}
