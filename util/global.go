package util

import "sync"

var (
	jwt         string
	serviceName string
	mu          sync.Mutex
)

func SetJWT(value string) {
	mu.Lock()
	defer mu.Unlock()
	jwt = value
}

func GetJWT() string {
	mu.Lock()
	defer mu.Unlock()
	return jwt
}

func SetServiceName(value string) {
	mu.Lock()
	defer mu.Unlock()
	serviceName = value
}

func GetServiceName() string {
	mu.Lock()
	defer mu.Unlock()
	return serviceName
}
