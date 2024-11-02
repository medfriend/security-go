package module

import (
	"gorm.io/gorm"
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
)

func InitializeAuthModule(db *gorm.DB) *controller.AuthController {
	userRolRepo := repository.NewUserRolRepository(db)
	userRepo := repository.NewUserRepository(db)

	userRolService := service.NewUserRolService(userRolRepo)
	userService := service.NewUserService(userRepo)

	authService := service.NewAuthService(
		userRolService,
		userService,
	)

	return controller.NewAuthController(authService)
}
