package repository

type FindUserFunc = func(email string, password string) (user *User, err error)

type User struct {
	_Id      string
	Email    string
	Password string
}
