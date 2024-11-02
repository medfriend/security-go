// filename: auth.module.go
// go:build wireinject
//go:build wireinject
// +build wireinject

package module

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
)

var AuthSet = wire.NewSet(
	repository.NewUserRepository,
	repository.NewUserRolRepository,
	repository.NewRoleResourceRepository,
	service.NewUserService,
	service.NewUserRolService,
	service.NewRoleResourceService,
	service.NewAuthService,
	controller.NewAuthController,
)

func InitializeAuthModule(db *gorm.DB) *controller.AuthController {
	wire.Build(AuthSet)
	return nil
}
