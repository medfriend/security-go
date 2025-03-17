package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"security-go/module"
)

func NewEntityRouter(router *gin.RouterGroup, db *gorm.DB) {

	EntityController := module.InitializeEntityModule(db)

	routerGroup := router.Group("entity")

	routerGroup.POST("/createEntity", EntityController.CreateEntity)
	routerGroup.GET("/getEntityById/:id", EntityController.GetEntityById)
	routerGroup.PUT("/updateEntity/:id", EntityController.UpdateEntity)
	routerGroup.DELETE("/deleteEntity/:id", EntityController.DeleteEntity)
	routerGroup.GET("/all", EntityController.GetAllEntities)
}

func init() {
	RegisterRouter(NewEntityRouter)
}
