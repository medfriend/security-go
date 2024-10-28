package module

import (
	"gorm.io/gorm"
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
	"security-go/util"
)

func InitializeMenuModule(db *gorm.DB) *controller.MenuController {
	return util.GenericModuleInitializer(
		db,
		repository.NewMenuRepository,
		service.NewMenuService,
		controller.NewMenuController,
	)
}
