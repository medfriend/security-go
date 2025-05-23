package entity

import "time"

// TableName especifica el nombre de la tabla
func (TareaProgramada) TableName() string {
	return "public.tareas_programadas"
}

type TareaProgramada struct {
	TareasProgramadasID   uint      `gorm:"column:tareas_programadas_id;primaryKey" json:"tareasProgramadasId"`
	Nombre                string    `gorm:"column:nombre" json:"nombre"`
	Descripcion           string    `gorm:"column:descripcion" json:"descripcion"`
	ExpresionProgramacion string    `gorm:"column:expresion_programacion" json:"expresionProgramacion"`
	TipoAccion            string    `gorm:"column:tipo_accion" json:"tipoAccion"`
	DatosAccion           string    `gorm:"column:datos_accion" json:"datosAccion"`
	Activo                bool      `gorm:"column:activo" json:"activo"`
	FechaCreacion         time.Time `gorm:"column:fecha_creacion" json:"fechaCreacion"`
	FechaActualizacion    time.Time `gorm:"column:fecha_actualizacion" json:"fechaActualizacion"`
	TiempoMaximo          int       `gorm:"column:tiempo_maximo" json:"tiempoMaximo"`
	PoliticaReintentos    int       `gorm:"column:politica_reintentos" json:"politicaReintentos"`
	MdServiceServicioID   uint      `gorm:"column:md_service_servicio_id" json:"mdServiceServicioId"`
}
