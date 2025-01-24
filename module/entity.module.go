// filename: entity.module.go
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

var EntitySet = wire.NewSet(
	repository.NewEntityRepository,
	service.NewEntityService,
	controller.NewEntityController,
)

func InitializeEntityModule(db *gorm.DB) *controller.EntityController {
	wire.Build(EntitySet)
	return nil
}
