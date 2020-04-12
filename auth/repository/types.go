package repository

type FindUserFunc = func(email string, password string) (user *User, err error)

type User struct {
	Email string
}
