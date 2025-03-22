//go:build wireinject
// +build wireinject

// filename: trazabilidad_usuario_accion.module.go
// go:build wireinject
package module

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
)

var trazaSet = wire.NewSet(
	repository.NewTrazabilidadUsuarioAccionRepository,
	service.NewTrazabilidadUsuarioAccionService,
	controller.NewTrazabilidadUsuarioAccionController,
)

func InitializeTrazabilidadUsuarioAccionModule(db *gorm.DB) *controller.TrazabilidadUsuarioAccionController {
	wire.Build(trazaSet)
	return nil
}
