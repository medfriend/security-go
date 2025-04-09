package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"security-go/module"
)

func NewUserRouter(router *gin.RouterGroup, db *gorm.DB) {

	userController := module.InitializeUserModule(db)

	routerGroup := router.Group("user")

	routerGroup.POST("/createuser", userController.CreateUser)
	routerGroup.GET("/byId/:id/", userController.GetUserById)
	routerGroup.GET("/all", userController.GetUsers)
	routerGroup.PUT("/update", userController.UpdateUser)
	routerGroup.DELETE("/deleteuser/:id", userController.DeleteUser)
	routerGroup.GET("filterUser/:query", userController.FilterUser)
}

func init() {
	RegisterRouter(NewUserRouter)
}
