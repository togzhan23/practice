package auth

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Authenticate(username, password string) bool {
	return username == "admin" && password == "password"
}
