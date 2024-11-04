package httpServer

import (
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "security-go/docs"
	"security-go/middleware"
	"security-go/router"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitHttpServer(taskQueue chan *http.Request, db *gorm.DB, serviceInfo map[string]string) {
	r := gin.Default()

	r.Use(middleware.PostResponseMiddleware())

	api := r.Group(serviceInfo["SERVICE_NAME"])

	fmt.Println(taskQueue)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.InitializeAllRouters(api, db)

	err := r.Run(fmt.Sprintf(":%s", serviceInfo["SERVICE_PORT"]))

	if err != nil {
		return
	}
}
