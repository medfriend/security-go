package module

import (
	"gorm.io/gorm"
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
	"security-go/util"
)

func InitializeUserRolModule(db *gorm.DB) *controller.UserRolController {
	return util.GenericModuleInitializer(
		db,
		repository.NewUserRolRepository,
		service.NewUserRolService,
		controller.NewUserRolController,
	)
}
