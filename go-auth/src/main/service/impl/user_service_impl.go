package impl

import (
	"goauth/src/main/service"
)

type UserServiceImpl struct {
}

func (u *UserServiceImpl) CreateUser() {
}

func (u *UserServiceImpl) CreateRole() {
}

func (u *UserServiceImpl) AssignRole() {
}

func NewUserServiceImpl() service.IUserService {
	return &UserServiceImpl{}
}
