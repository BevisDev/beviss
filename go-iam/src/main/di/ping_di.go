//go:build wireinject

package di

import (
	"github.com/google/wire"
	"goiam/src/main/controller"
	"goiam/src/main/repository/repositoryImpl"
	"goiam/src/main/service/impl"
)

func NewPingDI() *controller.PingController {
	wire.Build(
		impl.NewPingServiceImpl,
		repositoryImpl.NewPingRepositoryImpl,
		controller.NewPingController,
	)
	return new(controller.PingController)
}
