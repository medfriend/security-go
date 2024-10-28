package httpServer

import (
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "security-go/docs"

	"net/http"
	"security-go/router"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitHttpServer(taskQueue chan *http.Request, db *gorm.DB, serviceInfo map[string]string) {
	r := gin.Default()
	api := r.Group(serviceInfo["SERVICE_NAME"])

	fmt.Println(taskQueue)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.NewUserRouter(api, db)
	router.NewEntityRouter(api, db)
	router.NewPermisoRouter(api, db)
	router.NewResourceRouter(api, db)
	router.NewMenuRouter(api, db)

	err := r.Run(fmt.Sprintf(":%s", serviceInfo["SERVICE_PORT"]))

	if err != nil {
		return
	}
}
