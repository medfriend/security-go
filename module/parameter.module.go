// filename: parameter.module.go
// go:build wireinject
//go:build wireinject
// +build wireinject

package module

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
)

var parameterSet = wire.NewSet(
	repository.NewParameterRepository,
	service.NewParameterService,
	controller.NewParameterController,
)

func InitializeParameterModule(db *gorm.DB) *controller.ParameterController {
	wire.Build(parameterSet)
	return nil
}

func InitializeParameterService(db *gorm.DB) service.ParameterService {
	wire.Build(parameterSet)
	return nil
}
