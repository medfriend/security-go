// filename: permission.module.go
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

var PermisoSet = wire.NewSet(
	repository.NewPermisoRepository,
	service.NewPermisoService,
	controller.NewPermisoController,
)

func InitializePermisoModule(db *gorm.DB) *controller.PermisoController {
	wire.Build(PermisoSet)
	return nil
}
