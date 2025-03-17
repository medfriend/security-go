package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"security-go/module"
)

func NewPermisoRouter(router *gin.RouterGroup, db *gorm.DB) {

	permisoController := module.InitializePermisoModule(db)

	routerGroup := router.Group("permission")

	routerGroup.POST("/createPermiso", permisoController.CreatePermiso)
	routerGroup.GET("/getPermisoById/:id", permisoController.GetPermisoById)
	routerGroup.PUT("/updatePermiso/:id", permisoController.UpdatePermiso)
	routerGroup.DELETE("/deletePermiso/:id", permisoController.DeletePermiso)
}

func init() {
	RegisterRouter(NewPermisoRouter)
}
