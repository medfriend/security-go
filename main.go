package main

import (
	"fmt"
	"github.com/medfriend/shared-commons-go/util/consul"
	"github.com/medfriend/shared-commons-go/util/env"
	"github.com/medfriend/shared-commons-go/util/worker"
	"net/http"
	"runtime"
	"security-go/httpServer"
)

func main() {
	env.LoadEnv()

	consulClient := consul.ConnectToConsul()

	numCPUs := runtime.NumCPU()

	fmt.Printf("Detected %d CPUs, creating %d workers\n", numCPUs, numCPUs)

	taskQueue := make(chan *http.Request, 100)

	stop := make(chan struct{})

	worker.CreateWorkers(numCPUs, stop, taskQueue)

	go httpServer.InitHttpServer(taskQueue)

	worker.HandleShutdown(stop, consulClient)
}
