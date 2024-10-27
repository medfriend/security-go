package httpServer

import (
	"fmt"
	"net/http"
	"os"
	"security-go/router"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitHttpServer(taskQueue chan *http.Request, db *gorm.DB) {
	r := gin.Default()
	api := r.Group(os.Getenv("SERVICE_PATH"))

	fmt.Println(taskQueue)

	router.NewUserRouter(api, db)
	router.NewEntityRouter(api, db)
	router.NewPermisoRouter(api, db)
	router.NewResourceRouter(api, db)

	err := r.Run(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")))

	if err != nil {
		return
	}
}
