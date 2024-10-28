package main

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/medfriend/shared-commons-go/util/consul"
	"github.com/medfriend/shared-commons-go/util/env"
	gormUtil "github.com/medfriend/shared-commons-go/util/gorm"
	"github.com/medfriend/shared-commons-go/util/worker"
	"gorm.io/gorm"
	"net/http"
	"os"
	"runtime"
	"security-go/httpServer"
)

var db *gorm.DB

func handlerServiceInfo(consulClient *api.Client) map[string]string {

	serviceInfo, _ := consul.GetKeyValue(consulClient, os.Getenv("SERVICE_NAME"))
	var result map[string]string
	json.Unmarshal([]byte(serviceInfo), &result)

	return result
}

func main() {
	env.LoadEnv()

	consulClient := consul.ConnectToConsulKey(os.Getenv("SERVICE_NAME"))

	serviveInfo := handlerServiceInfo(consulClient)

	numCPUs := runtime.NumCPU()

	fmt.Printf("Detected %d CPUs, creating %d workers\n", numCPUs, numCPUs)

	taskQueue := make(chan *http.Request, 100)

	stop := make(chan struct{})

	worker.CreateWorkers(numCPUs, stop, taskQueue)

	initDB, err := gormUtil.InitDB(db, consulClient)

	if err != nil {
		return
	}

	httpServer.InitHttpServer(taskQueue, initDB, serviveInfo)

	worker.HandleShutdown(stop, consulClient)
}
