// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package module

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
)

// Injectors from auth.module.go:

func InitializeAuthModule(db *gorm.DB) *controller.AuthController {
	userRolRepository := repository.NewUserRolRepository(db)
	userRolService := service.NewUserRolService(userRolRepository)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	roleResourceRepository := repository.NewRoleResourceRepository(db)
	roleResourceService := service.NewRoleResourceService(roleResourceRepository)
	menuRepository := repository.NewMenuRepository(db)
	menuService := service.NewMenuService(menuRepository)
	resourcePermissionRepository := repository.NewResourcePermissionRepository(db)
	resourcePermissionService := service.NewResourcePermissionService(resourcePermissionRepository)
	authService := service.NewAuthService(userRolService, userService, roleResourceService, menuService, resourcePermissionService)
	authController := controller.NewAuthController(authService)
	return authController
}

// Injectors from entity.module.go:

func InitializeEntityModule(db *gorm.DB) *controller.EntityController {
	entityRepository := repository.NewEntityRepository(db)
	entityService := service.NewEntityService(entityRepository)
	entityController := controller.NewEntityController(entityService)
	return entityController
}

// Injectors from menu.module.go:

func InitializeMenuModule(db *gorm.DB) *controller.MenuController {
	menuRepository := repository.NewMenuRepository(db)
	menuService := service.NewMenuService(menuRepository)
	menuController := controller.NewMenuController(menuService)
	return menuController
}

// Injectors from permission.module.go:

func InitializePermisoModule(db *gorm.DB) *controller.PermisoController {
	permisoRepository := repository.NewPermisoRepository(db)
	permisoService := service.NewPermisoService(permisoRepository)
	permisoController := controller.NewPermisoController(permisoService)
	return permisoController
}

// Injectors from resource.module.go:

func InitializeResourceModule(db *gorm.DB) *controller.ResourceController {
	resourceRepository := repository.NewResourceRepository(db)
	resourceService := service.NewResourceService(resourceRepository)
	resourceController := controller.NewResourceController(resourceService)
	return resourceController
}

// Injectors from resource_permission.module.go:

func InitializeResourcePermissionModule(db *gorm.DB) *controller.ResourcePermissionController {
	resourcePermissionRepository := repository.NewResourcePermissionRepository(db)
	resourcePermissionService := service.NewResourcePermissionService(resourcePermissionRepository)
	resourcePermissionController := controller.NewResourcePermissionController(resourcePermissionService)
	return resourcePermissionController
}

// Injectors from rol.module.go:

func InitializeRolModule(db *gorm.DB) *controller.RolController {
	rolRepository := repository.NewRolRepository(db)
	rolService := service.NewRolService(rolRepository)
	rolController := controller.NewRolController(rolService)
	return rolController
}

// Injectors from role_resource.module.go:

func InitializeRoleResourceModule(db *gorm.DB) *controller.RoleResourceController {
	roleResourceRepository := repository.NewRoleResourceRepository(db)
	roleResourceService := service.NewRoleResourceService(roleResourceRepository)
	roleResourceController := controller.NewRoleResourceController(roleResourceService)
	return roleResourceController
}

// Injectors from user.module.go:

func InitializeUserModule(db *gorm.DB) *controller.UserController {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	return userController
}

// Injectors from userRol.module.go:

func InitializeUserRolModule(db *gorm.DB) *controller.UserRolController {
	userRolRepository := repository.NewUserRolRepository(db)
	userRolService := service.NewUserRolService(userRolRepository)
	userRolController := controller.NewUserRolController(userRolService)
	return userRolController
}

// auth.module.go:

var AuthSet = wire.NewSet(repository.NewResourcePermissionRepository, repository.NewUserRepository, repository.NewUserRolRepository, repository.NewRoleResourceRepository, repository.NewMenuRepository, service.NewResourcePermissionService, service.NewUserService, service.NewUserRolService, service.NewRoleResourceService, service.NewAuthService, service.NewMenuService, controller.NewAuthController)

// entity.module.go:

var EntitySet = wire.NewSet(repository.NewEntityRepository, service.NewEntityService, controller.NewEntityController)

// menu.module.go:

var MenuSet = wire.NewSet(repository.NewMenuRepository, service.NewMenuService, controller.NewMenuController)

// permission.module.go:

var PermisoSet = wire.NewSet(repository.NewPermisoRepository, service.NewPermisoService, controller.NewPermisoController)

// resource.module.go:

var ResourceSet = wire.NewSet(repository.NewResourceRepository, service.NewResourceService, controller.NewResourceController)

// resource_permission.module.go:

var resourcePermissionSet = wire.NewSet(repository.NewResourcePermissionRepository, service.NewResourcePermissionService, controller.NewResourcePermissionController)

// rol.module.go:

var RolSet = wire.NewSet(repository.NewRolRepository, service.NewRolService, controller.NewRolController)

// role_resource.module.go:

var RoleResourceSet = wire.NewSet(repository.NewRoleResourceRepository, service.NewRoleResourceService, controller.NewRoleResourceController)

// user.module.go:

var UserSet = wire.NewSet(repository.NewUserRepository, service.NewUserService, controller.NewUserController)

// userRol.module.go:

var UserRoleSet = wire.NewSet(repository.NewUserRolRepository, service.NewUserRolService, controller.NewUserRolController)
