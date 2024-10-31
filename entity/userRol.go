package entity

import "time"

func (UserRol) TableName() string {
	return "re-usuario-rol"
}

type UserRol struct {
	UsuarioRolID uint      `gorm:"primaryKey;autoIncrement;column:re-usuario-rol_id"`
	UsuarioID    *uint     `gorm:"column:usuario_id" json:"usuario_id"`
	RolID        *uint     `gorm:"column:rol_id" json:"rol_id"`
	EntidadID    *uint     `gorm:"column:entidad_id" json:"entidad_id"`
	FechaInicio  time.Time `gorm:"column:fecha_inicio" json:"fecha_inicio"`
	FechaFin     time.Time `gorm:"column:fecha_fin" json:"fecha_fin"`
	Indefinido   string    `gorm:"column:indefinido" json:"indefinido"`
}
