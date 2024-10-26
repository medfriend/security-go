package module

import (
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
	"security-go/util"

	"gorm.io/gorm"
)

func InitializeResourceModule(db *gorm.DB) *controller.ResourceController {
	return util.GenericModuleInitializer(
		db,
		repository.NewResourceRepository,
		service.NewResourceService,
		controller.NewResourceController,
	)
}