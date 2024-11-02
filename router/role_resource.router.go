package router

import (
	"security-go/module"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRoleResourceRouter(router *gin.RouterGroup, db *gorm.DB) {

	RoleResourceController := module.InitializeRoleResourceModule(db)

	routerGroup := router.Group("rol-recurso")

	routerGroup.POST("/asignar", RoleResourceController.CreateRoleResource)
	routerGroup.GET("/:id", RoleResourceController.GetRoleResourceById)
	routerGroup.PUT("/:id", RoleResourceController.UpdateRoleResource)
	routerGroup.DELETE("/desasignar/:id", RoleResourceController.DeleteRoleResource)

}
