package module

import (
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
	"security-go/util"

	"gorm.io/gorm"
)

func InitializeResourcePermissionModule(db *gorm.DB) *controller.ResourcePermissionController {
	return util.GenericModuleInitializer(
		db,
		repository.NewResourcePermissionRepository,
		service.NewResourcePermissionService,
		controller.NewResourcePermissionController,
	)
}