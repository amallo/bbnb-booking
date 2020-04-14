package usecase

import (
	repository "bbnb-booking/auth/repository"
	session "bbnb-booking/auth/session"
	"bbnb-booking/models"
)

func SignIn(findUser repository.FindUserFunc, signSession session.SignSessionFunc) AuthFunc {
	return func(credentials Credentials) (*string, *models.User, error) {
		user, err := findUser(credentials.Email, credentials.Password)
		if err != nil {
			return nil, nil, err
		}
		// Token should be enough to identify user with their email and id
		tokenResult, err := signSession(user.Email)
		if err != nil {
			return nil, nil, err
		}
		// We NEVER return password here
		userResult := models.User{Email: user.Email}
		return tokenResult, &userResult, nil
	}
}
