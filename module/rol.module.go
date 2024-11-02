// filename: rol.module.go
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

var RolSet = wire.NewSet(
	repository.NewRolRepository,
	service.NewRolService,
	controller.NewRolController,
)

func InitializeRolModule(db *gorm.DB) *controller.RolController {
	wire.Build(RolSet)
	return nil
}
