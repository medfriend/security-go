// filename: resource_permission.module.go
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

var resourcePermissionSet = wire.NewSet(
	repository.NewResourcePermissionRepository,
	service.NewResourcePermissionService,
	controller.NewResourcePermissionController,
)

func InitializeResourcePermissionModule(db *gorm.DB) *controller.ResourcePermissionController {
	wire.Build(resourcePermissionSet)
	return nil
}
