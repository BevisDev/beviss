package service

import (
	request2 "gobaucua/src/main/dto/request"
)

type IAuthService interface {
	SignIn(dto request2.SignInDTO)
	SignUp(dto request2.SignUpDTO)
}
