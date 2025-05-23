package scheduler

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"security-go/repository"
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

// CreateSchedule obtiene la informacion de las tareas registradas en la base de datos y las relaciona con las tareas registradas en la programacion
func CreateSchedule(db *gorm.DB) {
	tareaProgramadasRepo := repository.NewTareaProgramadaRepository(db)

	tareas, err := tareaProgramadasRepo.Find()

	fmt.Println(tareas, err)
}
