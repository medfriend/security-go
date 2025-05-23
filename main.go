package main

// @title           medfri-security
// @version         1.0
// @description     micro de seguridad.

// @host            localhost:9000
// @BasePath        /medfri-security

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Ingresa "Bearer {token}" para autenticar.

// @contact.name    Soporte de API
// @contact.url     http://www.soporte-api.com
// @contact.email   soporte@api.com

// @license.name    MIT
// @license.url     https://opensource.org/licenses/MIT

import (
	"fmt"
	"github.com/medfriend/shared-commons-go/util/consul"
	"github.com/medfriend/shared-commons-go/util/env"
	gormUtil "github.com/medfriend/shared-commons-go/util/gorm"
	"github.com/medfriend/shared-commons-go/util/migrations"
	"github.com/medfriend/shared-commons-go/util/worker"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"security-go/httpServer"
	"security-go/scheduler"
	"security-go/util"
	"syscall"
)

var db *gorm.DB

func main() {
	env.LoadEnv()

	consulIp := os.Getenv("CONSUL_IP")
	consulConn := fmt.Sprint(consulIp, ":8500")

	consulClient := consul.ConnectToConsulKey(consulConn, "SECURITY")

	serviceInfo := util.HandlerServiceInfo(consulClient)

	// Configuración del pool de workers usando la biblioteca compartida
	numCPUs := runtime.NumCPU()
	log.Printf("Detectados %d CPUs, creando un pool de workers con %d workers", numCPUs, numCPUs)

	// Crear un pool con capacidad para 100 solicitudes en cola
	workerPool := worker.NewWorkerPool(numCPUs, 100)

	// Configurar el worker pool como global para middleware
	worker.SetGlobalWorkerPool(workerPool)

	// Canal para señales de cierre
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// Inicialización de la base de datos
	initDB, err := gormUtil.InitDB(
		db,
		consulClient,
		"LOCAL",
		"SECURITY",
	)

	if err != nil {
		log.Fatalf("Error inicializando la base de datos: %v", err)
	}

	// Aplicar migraciones
	err = migrations.ReadMigration(initDB)
	if err != nil {
		log.Fatalf("Error aplicando migraciones: %v", err)
	}

	// Iniciar los scheduler programados
	scheduler.CreateSchedule(initDB)

	// Para mantener compatibilidad con el código existente
	legacyTaskQueue := make(chan *http.Request, 100)
	legacyStopChan := make(chan struct{})

	worker.CreateWorkers(numCPUs, legacyStopChan, legacyTaskQueue)
	go worker.HandleShutdown(legacyStopChan, consulClient)

	// Iniciar el servidor HTTP con nuestro nuevo workerPool
	serverChan := make(chan error, 1)

	go func() {
		if err := httpServer.InitHttpServerWithWorkerPool(workerPool, initDB, serviceInfo); err != nil {
			serverChan <- err
		}
	}()

	// Manejar señales de cierre
	select {
	case <-stopChan:
		log.Println("Señal de cierre recibida, apagando servicios...")
		// Cerrar el pool de workers
		workerPool.Shutdown()
		// Cerrar el canal legacy
		close(legacyStopChan)
	case err := <-serverChan:
		log.Printf("Error en el servidor HTTP: %v", err)
	}
}
