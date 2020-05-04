package services

import (
	"simple-auth/src/requests"
	"simple-auth/src/responses"
)

//AuthService godoc
type AuthService struct {
}

//Login godoc
func (authService *AuthService) Login(model requests.Login) responses.Login {
	return responses.Login{}
}
