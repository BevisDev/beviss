//go:build wireinject

package di

import (
	"github.com/google/wire"
	"goauth/src/main/controller"
	"goauth/src/main/repository/repositoryImpl"
	"goauth/src/main/service/impl"
)

func NewPingDI() *controller.PingController {
	wire.Build(
		impl.NewPingServiceImpl,
		repositoryImpl.NewPingRepositoryImpl,
		controller.NewPingController,
	)
	return new(controller.PingController)
}
