package entity

func (Permiso) TableName() string {
	return "permiso"
}

type Permiso struct {
	PermisoID uint   `gorm:"primaryKey;autoIncrement" json:"permiso_id"`
	Nombre    string `gorm:"size:128;not null" json:"nombre"`
	Accion    string `gorm:"size:2;not null" json:"accion"`
}
