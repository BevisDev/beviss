//go:build wireinject

package di

import (
	"github.com/google/wire"
	"gobaucua/src/main/controller"
	"gobaucua/src/main/repository/repositoryImpl"
	"gobaucua/src/main/service/impl"
)

func NewSupportDI() *controller.SupportController {
	wire.Build(
		impl.NewSupportServiceImpl,
		repositoryImpl.NewSupportRepositoryImpl,
		controller.NewSupportController,
	)
	return new(controller.SupportController)
}
