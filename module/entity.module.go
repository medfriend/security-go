package module

import (
	"gorm.io/gorm"
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
	"security-go/util"
)

func InitializeEntityModule(db *gorm.DB) *controller.EntityController {
	return util.GenericModuleInitializer(
		db,
		repository.NewEntityRepository,
		service.NewEntityService,
		controller.NewEntityController,
	)
}
