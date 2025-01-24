// filename: userRol.module.go
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

var UserRoleSet = wire.NewSet(
	repository.NewUserRolRepository,
	service.NewUserRolService,
	controller.NewUserRolController,
)

func InitializeUserRolModule(db *gorm.DB) *controller.UserRolController {
	wire.Build(UserRoleSet)
	return nil
}
