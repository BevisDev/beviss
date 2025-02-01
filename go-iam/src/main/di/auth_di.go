//go:build wireinject

package di

import (
	"github.com/google/wire"
	"goiam/src/main/controller"
	"goiam/src/main/service/impl"
)

func NewAuthDI() *controller.AuthController {
	wire.Build(
		impl.NewAuthServiceImpl,
		controller.NewAuthController,
	)
	return new(controller.AuthController)
}
