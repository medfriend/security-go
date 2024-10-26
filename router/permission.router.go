package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"security-go/module"
)

func NewPermisoRouter(router *gin.RouterGroup, db *gorm.DB) {

	permisoController := module.InitializePermisoModule(db)

	routerGroup := router.Group("permission")

	routerGroup.POST("/", permisoController.CreatePermiso)
	routerGroup.GET("/:id", permisoController.GetPermisoById)
	routerGroup.PUT("/:id", permisoController.UpdatePermiso)
	routerGroup.DELETE("/:id", permisoController.DeletePermiso)
}
