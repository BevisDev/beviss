package impl

import (
	request2 "goauth/src/main/dto/request"
	"goauth/src/main/service"
)

type AuthServiceImpl struct {
}

func (a *AuthServiceImpl) SignUp(dto request2.SignUpDTO) {

}

func (a *AuthServiceImpl) SignIn(dto request2.SignInDTO) {
}

func NewAuthServiceImpl() service.IAuthService {
	return &AuthServiceImpl{}
}
