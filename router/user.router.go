package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"security-go/module"
)

func NewUserRouter(router *gin.RouterGroup, db *gorm.DB) {

	userController := module.InitializeUserModule(db)

	routerGroup := router.Group("user")

	routerGroup.POST("/", userController.CreateUser)
	routerGroup.GET("/byId/:id/", userController.GetUserById)
	routerGroup.GET("/all", userController.GetUsers)
	routerGroup.PUT("/:id", userController.UpdateUser)
	routerGroup.DELETE("/:id", userController.DeleteUser)
}

func init() {
	RegisterRouter(NewUserRouter)
}
