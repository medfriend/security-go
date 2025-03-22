package entity

import "time"

type TrazabilidadUsuarioAccion struct {
	TrazabilidadUsuarioAccionID uint      `gorm:"primaryKey;column:trazabalidad_usuario_accion_id" json:"trazabalidad_usuario_accion_id"`
	UsuarioID                   uint      `gorm:"column:usuario_id" json:"usuario_id"`
	Accion                      string    `gorm:"column:accion" json:"accion"`
	FechaEjecucion              time.Time `gorm:"column:fecha_ejecucion;default:CURRENT_TIMESTAMP" json:"fecha_ejecucion"`
	Estado                      int       `gorm:"column:estado" json:"estado"`
	Endpoint                    string    `gorm:"column:endpoint" json:"endpoint"`
	Payload                     string    `gorm:"column:payload" json:"payload"`
	IP                          string    `gorm:"column:ip" json:"ip"`
	Error                       string    `gorm:"column:error" json:"error"`
	Duracion                    string    `gorm:"column:duracion" json:"duracion"`
	Coleccion                   string    `gorm:"column:coleccion" json:"coleccion"`
	Microservicio               string    `gorm:"column:microservicio" json:"microservicio"`
	CollectionID                string    `gorm:"column:collection_id" json:"collection_id"`
}

func (TrazabilidadUsuarioAccion) TableName() string { return "trazabilidad-usuario-accion" }
