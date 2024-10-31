package router

import (
	"security-go/module"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRoleResourceRouter(router *gin.RouterGroup, db *gorm.DB) {

	RoleResourceController := module.InitializeRoleResourceModule(db)

	routerGroup := router.Group("role_resource")

	routerGroup.POST("/", RoleResourceController.CreateRoleResource)
	routerGroup.GET("/:id", RoleResourceController.GetRoleResourceById)
	routerGroup.PUT("/:id", RoleResourceController.UpdateRoleResource)
	routerGroup.DELETE("/:id", RoleResourceController.DeleteRoleResource)

}