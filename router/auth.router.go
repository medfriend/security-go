package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"security-go/module"
)

func NewAuthRouter(router *gin.RouterGroup, db *gorm.DB) {
	AuthController := module.InitializeAuthModule(db)

	routerGroup := router.Group("auth")

	routerGroup.POST("/", AuthController.Login)
}
