//go:build wireinject
// +build wireinject

// filename: role_resource.module.go
// go:build wireinject
package module

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
)

var RoleResourceSet = wire.NewSet(
	repository.NewRoleResourceRepository,
	service.NewRoleResourceService,
	controller.NewRoleResourceController,
)

func InitializeRoleResourceModule(db *gorm.DB) *controller.RoleResourceController {
	wire.Build(RoleResourceSet)
	return nil
}
