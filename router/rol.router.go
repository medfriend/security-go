package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"security-go/module"
)

func NewRolRouter(router *gin.RouterGroup, db *gorm.DB) {

	RolController := module.InitializeRolModule(db)

	routerGroup := router.Group("rol")

	routerGroup.POST("/", RolController.CreateRol)
	routerGroup.GET("/:id", RolController.GetRolById)
	routerGroup.PUT("/:id", RolController.UpdateRol)
	routerGroup.DELETE("/:id", RolController.DeleteRol)

}

func init() {
	RegisterRouter(NewRolRouter)
}
