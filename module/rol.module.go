package module

import (
	"gorm.io/gorm"
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
	"security-go/util"
)

func InitializeRolModule(db *gorm.DB) *controller.RolController {
	return util.GenericModuleInitializer(
		db,
		repository.NewRolRepository,
		service.NewRolService,
		controller.NewRolController,
	)
}
