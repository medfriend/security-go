package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"security-go/module"
)

func NewMenuRouter(router *gin.RouterGroup, db *gorm.DB) {

	MenuController := module.InitializeMenuModule(db)

	routerGroup := router.Group("menu")

	routerGroup.POST("/", MenuController.CreateMenu)
	routerGroup.GET("/:id", MenuController.GetMenuById)
	routerGroup.PUT("/:id", MenuController.UpdateMenu)
	routerGroup.DELETE("/:id", MenuController.DeleteMenu)
	routerGroup.GET("/parents/:entidadId", MenuController.GetParentsMenuByEntity)
	routerGroup.GET("/childs-parent/:id", MenuController.GetChildByParentId)
	routerGroup.GET("/childs/:entidadId", MenuController.GetChildsMenuByEntity)
}

func init() {
	RegisterRouter(NewMenuRouter)
}
