package di

import (
	"github.com/google/wire"
	"goauth/src/main/controller"
	"goauth/src/main/service/impl"
)

func NewUserDI() *controller.UserController {
	wire.Build(
		impl.NewUserServiceImpl,
		controller.NewUserController,
	)
	return new(controller.UserController)
}
