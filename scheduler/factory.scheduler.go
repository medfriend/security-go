package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	mdfscheduler "github.com/medfriend/shared-commons-go/util/scheduler"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"security-go/entity"
	"security-go/module"
	"security-go/repository"
	"time"
)

func SetupFactory(serviceContainer *mdfscheduler.ServiceContainer) mdfscheduler.TaskFactory {

	factory := mdfscheduler.NewTaskFactory()
	parser := &mdfscheduler.DefaultDataParser{}

	// Registrar executor de servicios
	if serviceContainer != nil {
		factory.RegisterExecutor(
			mdfscheduler.TipoAccionService,
			mdfscheduler.NewServiceExecutor(parser, serviceContainer))
	}

	return factory
}

// CreateSchedule obtiene la informacion de las tareas registradas en la base de datos y las relaciona con las tareas registradas en la programacion
func CreateSchedule(db *gorm.DB) {
	// Crear contenedor de servicios y registrar todos los servicios
	serviceContainer := mdfscheduler.NewServiceContainer()

	serviceContainer.RegisterService("AuthService", module.InitilizeAuthService(db))
	serviceContainer.RegisterService("UserService", module.InitializeUserService(db))
	serviceContainer.RegisterService("RolService", module.InitializeRolService(db))
	serviceContainer.RegisterService("MenuService", module.InitializeMenuService(db))
	serviceContainer.RegisterService("PermissionService", module.InitializePermisoService(db))
	serviceContainer.RegisterService("ResourceService", module.InitializeResourceService(db))
	serviceContainer.RegisterService("ResourcePermissionService", module.InitializeResourcePermissionService(db))
	serviceContainer.RegisterService("RoleResourceService", module.InitializeRoleResourceService(db))
	serviceContainer.RegisterService("ParameterService", module.InitializeParameterService(db))
	serviceContainer.RegisterService("EntityService", module.InitializeEntityService(db))
	serviceContainer.RegisterService("UserRolService", module.InitializeUserRolService(db))
	serviceContainer.RegisterService("TrazabilidadService", module.InitializeTrazabilidadUsuarioAccionService(db))

	c := cron.New(cron.WithSeconds())

	tareaProgramadasRepo := repository.NewTareaProgramadaRepository(db)

	// Obtener todas las tareas activas
	tareas, err := tareaProgramadasRepo.Find()
	if err != nil {
		fmt.Printf("Error obteniendo tareas programadas: %v\n", err)
		return
	}

	// Configurar factory con todos los ejecutores
	factory := SetupFactory(serviceContainer)

	// Procesar cada tarea
	for _, tarea := range *tareas {
		if !tarea.Activo {
			fmt.Printf("Tarea '%s' está inactiva, omitiendo\n", tarea.Nombre)
			continue
		}

		// Obtener el ejecutor apropiado
		executor, err := factory.CreateExecutor(mdfscheduler.TipoAccion(tarea.TipoAccion))
		if err != nil {
			fmt.Printf("Error obteniendo ejecutor para tarea '%s': %v\n", tarea.Nombre, err)
			continue
		}

		// Convertir entity.TareaProgramada a scheduler.TareaProgramada
		schedulerTarea := convertToSchedulerTarea(tarea)

		c.AddFunc(tarea.ExpresionProgramacion, func() {
			// Ejecutar la tarea (ejemplo)
			ctx := context.Background()

			resultado, err := executor.Execute(ctx, schedulerTarea)

			if err != nil {
				fmt.Printf("Error ejecutando tarea '%s': %v\n", tarea.Nombre, err)
			} else {
				fmt.Printf("Tarea '%s' ejecutada exitosamente. Estado: %s\n",
					tarea.Nombre, resultado.Estado)

			}

			fmt.Printf("Programando tarea '%s' con expresión: %s\n",
				tarea.Nombre, tarea.ExpresionProgramacion)
		})
	}

	c.Start()

	fmt.Printf("Procesadas %d tareas programadas\n", len(*tareas))
}

// convertToSchedulerTarea convierte entity.TareaProgramada a scheduler.TareaProgramada
func convertToSchedulerTarea(entityTarea entity.TareaProgramada) *mdfscheduler.TareaProgramada {
	// Convertir DatosAccion de string a json.RawMessage
	var datosAccion json.RawMessage
	if entityTarea.DatosAccion != "" {
		datosAccion = json.RawMessage(entityTarea.DatosAccion)
	}

	// Convertir campos opcionales
	var tiempoMaximo *float64
	if entityTarea.TiempoMaximo > 0 {
		temp := float64(entityTarea.TiempoMaximo)
		tiempoMaximo = &temp
	}

	var politicaReintentos json.RawMessage
	if entityTarea.PoliticaReintentos > 0 {
		// Crear una política básica desde el entero
		politica := fmt.Sprintf(`{"max_reintentos": %d, "intervalo_reintento": "5m"}`, entityTarea.PoliticaReintentos)
		politicaReintentos = json.RawMessage(politica)
	}

	var servicioID *int
	if entityTarea.MdServiceServicioID > 0 {
		temp := int(entityTarea.MdServiceServicioID)
		servicioID = &temp
	}

	var fechaActualizacion *time.Time
	if !entityTarea.FechaActualizacion.IsZero() {
		fechaActualizacion = &entityTarea.FechaActualizacion
	}

	return &mdfscheduler.TareaProgramada{
		ID:                    int64(entityTarea.TareasProgramadasID),
		Nombre:                entityTarea.Nombre,
		Descripcion:           entityTarea.Descripcion,
		ExpresionProgramacion: entityTarea.ExpresionProgramacion,
		TipoAccion:            entityTarea.TipoAccion,
		DatosAccion:           datosAccion,
		Activo:                entityTarea.Activo,
		FechaCreacion:         entityTarea.FechaCreacion,
		FechaActualizacion:    fechaActualizacion,
		TiempoMaximo:          tiempoMaximo,
		PoliticaReintentos:    politicaReintentos,
		ServicioID:            servicioID,
	}
}
