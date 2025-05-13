package scheduler

import (
	"net/http"
	"time"
)

func SetupFactory() TaskFactory {
	factory := NewTaskFactory()
	parser := &DefaultDataParser{}

	// Registrar ejecutores
	httpClient := &http.Client{Timeout: 30 * time.Second}
	factory.RegisterExecutor(TipoAccionHTTP, NewHTTPExecutor(parser, httpClient))

	// Aquí registrarías otros ejecutores
	// factory.RegisterExecutor(models.TipoAccionScript, NewScriptExecutor(parser))

	return factory
}
