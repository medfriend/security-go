package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/medfriend/shared-commons-go/util/worker"
	"log"
)

// PostResponseMiddleware es un fallback para mantener compatibilidad con código antiguo
// Idealmente se debería usar el middleware worker.TrazabilidadMiddleware en su lugar
func PostResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Si estamos usando el worker pool compartido, delegar a su middleware
		if worker.GetGlobalWorkerPool() != nil {
			worker.TrazabilidadMiddleware(nil)(c)
			return
		}

		// Implementación legacy para el caso en que no haya worker pool configurado
		authorizationHeader := c.Request.Header.Get("Authorization")

		c.Next()

		if authorizationHeader != "" {
			go func() {
				log.Println("Ejecutando tarea de trazabilidad en segundo plano (modo legacy)...")
				// Lógica anterior de trazabilidad
				// rabbitMQ := util.GetInstance()
				// mensaje := fmt.Sprintf(`{"Authorization": "%s"}`, authorizationHeader)
				// rabbitMQ.SendMessage("trazabilidad-usuario-accion", mensaje)
			}()
		} else {
			log.Println("No se encontró el header Authorization en la solicitud.")
		}
	}
}
