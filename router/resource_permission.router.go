package router

import (
	"security-go/module"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewResourcePermissionRouter(router *gin.RouterGroup, db *gorm.DB) {
	ResourcePermissionController := module.InitializeResourcePermissionModule(db)

	routerGroup := router.Group("resource_permission")

	routerGroup.POST("/createResourcePermission", ResourcePermissionController.CreateResourcePermission)
	routerGroup.GET("/getResourcePermissionById/:id", ResourcePermissionController.GetResourcePermissionById)
	routerGroup.PUT("/updateResourcePermission/:id", ResourcePermissionController.UpdateResourcePermission)
	routerGroup.DELETE("/deleteResourcePermission/:id", ResourcePermissionController.DeleteResourcePermission)
}

func init() {
	RegisterRouter(NewResourcePermissionRouter)
}
