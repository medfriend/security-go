package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"security-go/module"
)

func NewRolRouter(router *gin.RouterGroup, db *gorm.DB) {

	RolController := module.InitializeRolModule(db)

	routerGroup := router.Group("rol")

	routerGroup.POST("/createRol", RolController.CreateRol)
	routerGroup.GET("/getRolById/:id", RolController.GetRolById)
	routerGroup.PUT("/updateRol/:id", RolController.UpdateRol)
	routerGroup.DELETE("/DeleteRol/:id", RolController.DeleteRol)
	routerGroup.GET("/all", RolController.GetRoles)

}

func init() {
	RegisterRouter(NewRolRouter)
}
