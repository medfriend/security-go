package main

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
		os.Getenv("SERVICE_STATUS"),
	)

	if err != nil {
		return
	}

	httpServer.InitHttpServer(taskQueue, initDB, serviceInfo)

	worker.HandleShutdown(stop, consulClient)
}
