// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"goiam/src/main/controller"
	"goiam/src/main/repository/repositoryImpl"
	"goiam/src/main/service/impl"
)

// Injectors from auth_di.go:

func NewAuthDI() *controller.AuthController {
	iAuthService := impl.NewAuthServiceImpl()
	authController := controller.NewAuthController(iAuthService)
	return authController
}

// Injectors from ping_di.go:

func NewPingDI() *controller.PingController {
	iPingRepository := repositoryImpl.NewPingRepositoryImpl()
	iPingService := impl.NewPingServiceImpl(iPingRepository)
	pingController := controller.NewPingController(iPingService)
	return pingController
}

// Injectors from user_di.go:

func NewUserDI() *controller.UserController {
	iUserService := impl.NewUserServiceImpl()
	userController := controller.NewUserController(iUserService)
	return userController
}
