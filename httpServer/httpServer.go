package httpServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"security-go/router"
)

func InitHttpServer(taskQueue chan *http.Request, db *gorm.DB) {
	r := gin.Default()
	api := r.Group(os.Getenv("SERVICE_PATH"))

	fmt.Println(taskQueue)

	router.NewUserRouter(api, db)
	router.NewHospitalRouter(api, db)

	err := r.Run(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")))

	if err != nil {
		return
	}
}
