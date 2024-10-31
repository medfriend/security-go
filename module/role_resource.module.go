package module

import (
	"gorm.io/gorm"
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
	"security-go/util"
)

func InitializeRoleResourceModule(db *gorm.DB) *controller.RoleResourceController {
	return util.GenericModuleInitializer(
		db,
		repository.NewRoleResourceRepository,
		service.NewRoleResourceService,
		controller.NewRoleResourceController,
	)
}