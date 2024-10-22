package httpServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func InitHttpServer(taskQueue chan *http.Request) {
	r := gin.Default()

	r.Any(
		fmt.Sprintf("%s/*path", os.Getenv("SERVICE_PATH")),
		func(c *gin.Context) {
			taskQueue <- c.Request
			c.JSON(200, gin.H{
				"data": gin.H{
					"status": "ok",
				},
			})
		})

	err := r.Run(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")))

	if err != nil {
		return
	}
}
