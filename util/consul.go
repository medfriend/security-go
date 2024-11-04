package util

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/medfriend/shared-commons-go/util/consul"
	"os"
)

func HandlerServiceInfo(consulClient *api.Client) map[string]string {

	serviceInfo, _ := consul.GetKeyValue(consulClient, os.Getenv("SERVICE_NAME"))
	rabbitInfo, _ := consul.GetKeyValue(consulClient, "RABBIT")

	jwt, _ := consul.GetKeyValue(consulClient, "JWT")

	var result map[string]string
	var resultRabbitmq map[string]string

	err := json.Unmarshal([]byte(serviceInfo), &result)

	if err != nil {
		return nil
	}

	err = json.Unmarshal([]byte(rabbitInfo), &resultRabbitmq)
	s := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		resultRabbitmq["RABBIT_USER"],
		resultRabbitmq["RABBIT_PASSWORD"],
		resultRabbitmq["RABBIT_HOST"],
		resultRabbitmq["RABBIT_PORT"])

	SetJWT(jwt)
	SetServiceName(result["SERVICE_NAME"])
	SetRabbitConn(s)

	return result
}
