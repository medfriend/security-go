package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"security-go/module"
)

func NewHospitalRouter(router *gin.RouterGroup, db *gorm.DB) {

	hospitalController := module.InitializeHospitalModule(db)

	routerGroup := router.Group("hospital")

	routerGroup.POST("/", hospitalController.CreateHospital)
	routerGroup.GET("/:id", hospitalController.GetHospitalById)
	routerGroup.PUT("/:id", hospitalController.UpdateHospital)
	routerGroup.DELETE("/:id", hospitalController.DeleteHospital)
}
