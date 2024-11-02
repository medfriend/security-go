//go:build wireinject
// +build wireinject

// filename: user.module.go
// go:build wireinject
package module

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
)

var UserSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
	controller.NewUserController,
)

func InitializeUserModule(db *gorm.DB) *controller.UserController {
	wire.Build(UserSet)
	return nil
}
