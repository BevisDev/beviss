//go:build wireinject

package di

import (
	"github.com/google/wire"
	"gobaucua/src/main/controller"
	"gobaucua/src/main/service/impl"
)

func NewUserDI() *controller.UserController {
	wire.Build(
		impl.NewUserServiceImpl,
		controller.NewUserController,
	)
	return new(controller.UserController)
}
