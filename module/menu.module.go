// filename: menu.module.go
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

var MenuSet = wire.NewSet(
	repository.NewMenuRepository,
	service.NewMenuService,
	controller.NewMenuController,
)

func InitializeMenuModule(db *gorm.DB) *controller.MenuController {
	wire.Build(MenuSet)
	return nil
}
