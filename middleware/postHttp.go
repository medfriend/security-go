package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"security-go/util"
)

// PostResponseMiddleware enviar la informacion a la traza de acciones
func PostResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Captura el header Authorization
		authorizationHeader := c.Request.Header.Get("Authorization")

		c.Next()

		go func() {
			log.Println("Ejecutando tarea de trazabilidad en segundo plano...") // Log para verificación

			if authorizationHeader != "" {

				rabbitMQ := util.GetInstance()
				mensaje := fmt.Sprintf(`{"Authorization": "%s"}`, authorizationHeader)

				rabbitMQ.SendMessage("trazabilidad-usuario-accion", mensaje)
			} else {
				log.Println("No se encontró el header Authorization en la solicitud.")
			}
		}()
	}
}
