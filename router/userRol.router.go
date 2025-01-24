package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"security-go/module"
)

func NewUserRolRouter(router *gin.RouterGroup, db *gorm.DB) {

	UserRolController := module.InitializeUserRolModule(db)

	routerGroup := router.Group("user-rol")

	routerGroup.POST("/asignar", UserRolController.CreateUserRol)
	routerGroup.DELETE("/desasignar/:id", UserRolController.DeleteUserRol)
}

func init() {
	RegisterRouter(NewUserRolRouter)
}
