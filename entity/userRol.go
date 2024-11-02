package entity

import "time"

func (UserRol) TableName() string {
	return "re-usuario-rol"
}

type UserRol struct {
	UsuarioRolID int64     `gorm:"primaryKey;autoIncrement;column:re-usuario-rol_id"`
	UsuarioID    *int64    `gorm:"column:usuario_id" json:"usuario_id"`
	RolID        *int64    `gorm:"column:rol_id" json:"rol_id"`
	EntidadID    *int64    `gorm:"column:entidad_id" json:"entidad_id"`
	FechaInicio  time.Time `gorm:"column:fecha_inicio" json:"fecha_inicio"`
	FechaFin     time.Time `gorm:"column:fecha_fin" json:"fecha_fin"`
	Indefinido   bool      `gorm:"column:indefinido" json:"indefinido"`
}
