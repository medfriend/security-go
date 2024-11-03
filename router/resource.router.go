package router

import (
	"security-go/module"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewResourceRouter(router *gin.RouterGroup, db *gorm.DB) {

	resourceController := module.InitializeResourceModule(db)

	routerGroup := router.Group("resources")

	routerGroup.POST("/", resourceController.CreateResource)
	routerGroup.GET("/:id", resourceController.GetResourceById)
	routerGroup.PUT("/:id", resourceController.UpdateResource)
	routerGroup.DELETE("/:id", resourceController.DeleteResource)
}

func init() {
	RegisterRouter(NewResourceRouter)
}
