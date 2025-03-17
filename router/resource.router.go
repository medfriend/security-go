package router

import (
	"security-go/module"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewResourceRouter(router *gin.RouterGroup, db *gorm.DB) {

	resourceController := module.InitializeResourceModule(db)

	routerGroup := router.Group("resources")

	routerGroup.POST("/createResource", resourceController.CreateResource)
	routerGroup.GET("/getResourceById/:id", resourceController.GetResourceById)
	routerGroup.PUT("/updateResource/:id", resourceController.UpdateResource)
	routerGroup.DELETE("/deleteResource/:id", resourceController.DeleteResource)
	routerGroup.GET("/all", resourceController.GetAllResources)
}

func init() {
	RegisterRouter(NewResourceRouter)
}
