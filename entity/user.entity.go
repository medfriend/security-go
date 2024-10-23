package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UsuarioID          uint       `gorm:"primaryKey;autoIncrement" json:"usuario_id"`
	Usuario            string     `gorm:"size:100;not null" json:"usuario"`
	Nombre1            string     `gorm:"size:100" json:"nombre_1"`
	Nombre2            string     `gorm:"size:100" json:"nombre_2"`
	ApellidoPaterno    string     `gorm:"size:100" json:"apellido_paterno"`
	ApellidoMaterno    string     `gorm:"size:100" json:"apellido_materno"`
	Clave              string     `gorm:"not null" json:"clave"`
	Email              string     `gorm:"size:100;unique;not null" json:"email"`
	CambiarClave       bool       `gorm:"default:false" json:"cambiar_clave"`
	FechaCambioClave   time.Time  `json:"fecha_cambio_clave"`
	FechaCreacion      time.Time  `gorm:"autoCreateTime" json:"fecha_creacion"`
	FechaRetiro        *time.Time `json:"fecha_retiro"`
	Activo             bool       `gorm:"default:true" json:"activo"`
	TiempoValidezToken int        `json:"tiempo_validez_token"`

	gorm.Model
}
