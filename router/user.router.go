package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"security-go/module"
)

func NewUserRouter(router *gin.RouterGroup, db *gorm.DB) {

	userController := module.InitializeUserModule(db)

	router.POST("/users", userController.CreateUser)
	router.GET("/users/:id", userController.GetUserById)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)
}
