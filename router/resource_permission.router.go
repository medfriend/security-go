package router

import (
	"security-go/module"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewResourcePermissionRouter(router *gin.RouterGroup, db *gorm.DB) {
	ResourcePermissionController := module.InitializeResourcePermissionModule(db)

	routerGroup := router.Group("resource_permission")

	routerGroup.POST("/", ResourcePermissionController.CreateResourcePermission)
	routerGroup.GET("/:id", ResourcePermissionController.GetResourcePermissionById)
	routerGroup.PUT("/:id", ResourcePermissionController.UpdateResourcePermission)
	routerGroup.DELETE("/:id", ResourcePermissionController.DeleteResourcePermission)
}

func init() {
	RegisterRouter(NewResourcePermissionRouter)
}
