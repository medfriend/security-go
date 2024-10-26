package module

import (
	"gorm.io/gorm"
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
	"security-go/util"
)

func InitializePermisoModule(db *gorm.DB) *controller.PermisoController {
	return util.GenericModuleInitializer(
		db,
		repository.NewPermisoRepository,
		service.NewPermisoService,
		controller.NewPermisoController,
	)
}
