//go:build wireinject

package di

import (
	"github.com/google/wire"
	"goiam/src/main/controller"
	"goiam/src/main/service/impl"
)

func NewUserDI() *controller.UserController {
	wire.Build(
		impl.NewUserServiceImpl,
		controller.NewUserController,
	)
	return new(controller.UserController)
}
