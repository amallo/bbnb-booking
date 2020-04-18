package usecase

import (
	repository "bbnb-booking/auth/repository"
	session "bbnb-booking/auth/session"
	"bbnb-booking/models"
	"fmt"
)

func SignInUseCase(
	findUser repository.FindUserFunc,
	signSession session.SignSessionFunc) AuthFunc {
	return func(credentials Credentials) (*string, *models.User, error) {
		criteria := repository.Criteria{"email": credentials.Email, "password": credentials.Password}
		user, err := findUser(criteria)
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
	findUser repository.FindUserFunc,
	insertUser repository.InsertUserFunc,
	signSession session.SignSessionFunc) AuthFunc {
	return func(credentials Credentials) (*string, *models.User, error) {
		_, err := findUser(repository.Criteria{"email": credentials.Email})
		if err != nil {
			if err.(*repository.RepositoryError).ErrNoDocuments() {
				user, err := insertUser(repository.Criteria{"email": credentials.Email, "password": credentials.Password})
				// Cannot insert, forward database error
				if err != nil {
					return nil, nil, fmt.Errorf("Cannot signup: %s", err.Error())
				}
				// Token should be enough to identify user with their  id
				tokenResult, err := signSession(user.ID.Hex())
				if err != nil {
					return nil, nil, fmt.Errorf("Cannot signup: %s", err.Error())
				}
				// We NEVER return password here
				userResult := models.User{Email: user.Email}
				return tokenResult, &userResult, nil
			}
		}
		return nil, nil, fmt.Errorf("Cannot signup: user already exists")
	}
}
