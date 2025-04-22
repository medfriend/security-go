package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"security-go/module"
)

func NewParameterRouter(router *gin.RouterGroup, db *gorm.DB) {
	ParameterController := module.InitializeParameterModule(db)

	routerGRoup := router.Group("parameter")

	routerGRoup.GET("/findParameterByCode/:code", ParameterController.FindParameterByCode)
}

func init() {
	RegisterRouter(NewParameterRouter)
}
