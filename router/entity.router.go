package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"security-go/module"
)

func NewEntityRouter(router *gin.RouterGroup, db *gorm.DB) {

	EntityController := module.InitializeEntityModule(db)

	routerGroup := router.Group("entity")

	routerGroup.POST("/", EntityController.CreateEntity)
	routerGroup.GET("/:id", EntityController.GetEntityById)
	routerGroup.PUT("/:id", EntityController.UpdateEntity)
	routerGroup.DELETE("/:id", EntityController.DeleteEntity)
}
