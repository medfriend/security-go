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
	"github.com/medfriend/shared-commons-go/util/worker"
	"gorm.io/gorm"
	"net/http"
	"os"
	"runtime"
	"security-go/httpServer"
	"security-go/util"
)

var db *gorm.DB

func main() {
	env.LoadEnv()

	consulClient := consul.ConnectToConsulKey(os.Getenv("SERVICE_NAME"))

	serviceInfo := util.HandlerServiceInfo(consulClient)

	numCPUs := runtime.NumCPU()

	fmt.Printf("Detected %d CPUs, creating %d workers\n", numCPUs, numCPUs)

	taskQueue := make(chan *http.Request, 100)

	stop := make(chan struct{})

	worker.CreateWorkers(numCPUs, stop, taskQueue)

	initDB, err := gormUtil.InitDB(
		db,
		consulClient,
		"LOCAL",
		"SECURITY",
	)

	if err != nil {
		return
	}

	httpServer.InitHttpServer(taskQueue, initDB, serviceInfo)

	worker.HandleShutdown(stop, consulClient)
}
