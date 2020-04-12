package usecase

type SignInCredentials struct {
	Email    string
	Password string
}

type SignInUseCase = func(SignInCredentials)

func SignIn(credentials SignInCredentials) {

}
