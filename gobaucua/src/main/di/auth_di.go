//go:build wireinject

package di

import (
	"github.com/google/wire"
	"gobaucua/src/main/controller"
	"gobaucua/src/main/service/impl"
)

func NewAuthDI() *controller.AuthController {
	wire.Build(
		impl.NewAuthServiceImpl,
		controller.NewAuthController,
	)
	return new(controller.AuthController)
}
