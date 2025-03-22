package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"security-go/module"
)

func NewTrazabilidadUsuarioAccion(router *gin.RouterGroup, db *gorm.DB) {
	trazaController := module.InitializeTrazabilidadUsuarioAccionModule(db)
	routerGroup := router.Group("trazabilidadUsuarioAccion")

	routerGroup.GET("/GetTrazaByUserId/:id", trazaController.GetTrazaByUserId)
}

func init() {
	RegisterRouter(NewTrazabilidadUsuarioAccion)
}
