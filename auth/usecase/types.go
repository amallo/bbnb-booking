package usecase

type Credentials struct {
	Email    string
	Password string
}

type SignInFunc = func(Credentials) (*string, error)
