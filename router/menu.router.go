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

}
