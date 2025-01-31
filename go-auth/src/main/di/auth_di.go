//go:build wireinject

package di

import (
	"github.com/google/wire"
	"goauth/src/main/controller"
	"goauth/src/main/service/impl"
)

func NewAuthDI() *controller.AuthController {
	wire.Build(
		impl.NewAuthServiceImpl,
		controller.NewAuthController,
	)
	return new(controller.AuthController)
}
