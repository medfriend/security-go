//go:build wireinject
// +build wireinject

// filename: resource.module.go
// go:build wireinject
package module

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
)

var ResourceSet = wire.NewSet(
	repository.NewResourceRepository,
	service.NewResourceService,
	controller.NewResourceController,
)

func InitializeResourceModule(db *gorm.DB) *controller.ResourceController {
	wire.Build(ResourceSet)
	return nil
}
