package usecase

import (
	repository "bbnb-booking/auth/repository"
	session "bbnb-booking/auth/session"
	"bbnb-booking/models"
	"fmt"
)

func SignInUseCase(
	findUserByEmailAndPassword repository.FindUserByEmailAndPasswordFunc,
	signSession session.SignSessionFunc) AuthFunc {
	return func(credentials Credentials) (*string, *models.User, error) {
		user, err := findUserByEmailAndPassword(credentials.Email, credentials.Password)
		if err != nil {
			return nil, nil, err
		}
		// Token should be enough to identify user with their  id
		tokenResult, err := signSession(user.ID.String())
		if err != nil {
			return nil, nil, err
		}
		// We NEVER return password here
		userResult := models.User{Email: user.Email}
		return tokenResult, &userResult, nil
	}
}

func SignUpUseCase(
	findUserByEmailAndPassword repository.FindUserByEmailAndPasswordFunc,
	insertUser repository.InsertUserFunc,
	signSession session.SignSessionFunc) AuthFunc {
	return func(credentials Credentials) (*string, *models.User, error) {
		_, err := findUserByEmailAndPassword(credentials.Email, credentials.Password)
		if err != nil {
			if err.(*repository.RepositoryError).ErrNoDocuments() {
				user, err := insertUser(credentials.Email, credentials.Password)
				if err != nil {
					return nil, nil, err
				}
				// Token should be enough to identify user with their  id
				tokenResult, err := signSession(user.ID.Hex())
				if err != nil {
					return nil, nil, err
				}
				// We NEVER return password here
				userResult := models.User{Email: user.Email}
				return tokenResult, &userResult, nil
			}
		}
		return nil, nil, fmt.Errorf("Cannot signup, user already exists")

	}
}
