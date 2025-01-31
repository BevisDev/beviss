// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	controller2 "goauth/src/main/controller"
	"goauth/src/main/repository/repositoryImpl"
	impl2 "goauth/src/main/service/impl"
)

// Injectors from auth_di.go:

func NewAuthDI() *controller2.AuthController {
	iAuthService := impl2.NewAuthServiceImpl()
	authController := controller2.NewAuthController(iAuthService)
	return authController
}

// Injectors from ping_di.go:

func NewPingDI() *controller2.PingController {
	iPingRepository := repositoryImpl.NewPingRepositoryImpl()
	iPingService := impl2.NewPingServiceImpl(iPingRepository)
	pingController := controller2.NewPingController(iPingService)
	return pingController
}

// Injectors from user_di.go:

func NewUserDI() *controller2.UserController {
	iUserService := impl2.NewUserServiceImpl()
	userController := controller2.NewUserController(iUserService)
	return userController
}
