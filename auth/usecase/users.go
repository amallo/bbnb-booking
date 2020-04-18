package usecase

import (
	repository "bbnb-booking/auth/repository"
	session "bbnb-booking/auth/session"

	"bbnb-booking/models"
	"fmt"
)

func SignInUseCase(
	findUser repository.FindUserFunc,
	md5Hash session.MD5HashFunc,
	signSession session.SignSessionFunc) AuthFunc {
	return func(credentials Credentials) (*string, *models.User, error) {
		// Hash passwords before requesting te repository
		hashPassword := md5Hash(credentials.Password)
		user, err := findUser(repository.Criteria{"email": credentials.Email, "password": hashPassword})
		if err != nil {
			return nil, nil, fmt.Errorf("Cannot signin: %s", err.Error())
		}
		// Token should be enough to identify user with their  id
		tokenResult, err := signSession(user.ID.String())
		if err != nil {
			return nil, nil, fmt.Errorf("Cannot signin: %s", err.Error())
		}
		// We NEVER return password here
		userResult := models.User{Email: user.Email}
		return tokenResult, &userResult, nil
	}
}

func SignUpUseCase(
	findUser repository.FindUserFunc,
	insertUser repository.InsertUserFunc,
	md5Hash session.MD5HashFunc,
	signSession session.SignSessionFunc) AuthFunc {
	return func(credentials Credentials) (*string, *models.User, error) {
		_, err := findUser(repository.Criteria{"email": credentials.Email})
		if err != nil {
			if err.(*repository.RepositoryError).ErrNoDocuments() {
				hashPassword := md5Hash(credentials.Password)
				user, err := insertUser(repository.Criteria{"email": credentials.Email, "password": hashPassword})
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
