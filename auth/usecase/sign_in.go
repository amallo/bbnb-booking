package usecase

import (
	repository "housings-api/auth/repository"
	session "housings-api/auth/session"
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
