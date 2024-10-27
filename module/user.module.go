package module

import (
	"gorm.io/gorm"
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
	"security-go/util"
)

func InitializeUserModule(db *gorm.DB) *controller.UserController {
	return util.GenericModuleInitializer(
		db,
		repository.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)
}
