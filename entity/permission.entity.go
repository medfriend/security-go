package entity

import "gorm.io/gorm"

type Permiso struct {
	PermisoID uint   `gorm:"primaryKey;autoIncrement" json:"permiso_id"`
	Nombre    string `gorm:"size:128;not null" json:"nombre"`
	Accion    string `gorm:"size:2;not null" json:"accion"`

	gorm.Model
}
