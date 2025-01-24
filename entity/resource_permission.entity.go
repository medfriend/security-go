package entity

func (ResourcePermission) TableName() string {
	return "re-recurso-permiso"
}

type ResourcePermission struct {
	RecursoPermisoID uint      `gorm:"primaryKey;autoIncrement;column:re-recurso-permiso_id"`
	RecursoID        *uint     `gorm:"column:recurso_id" json:"recurso_id,omitempty"`
	RolID            *uint     `gorm:"column:rol_id" json:"rol_id,omitempty"`
	PermisoID        *uint     `gorm:"column:permiso_id" json:"permiso_id,omitempty"`
	Recurso          *Resource `gorm:"foreignKey:RecursoID" json:"recurso,omitempty"`
	Rol              *Rol      `gorm:"foreignKey:RolID" json:"rol,omitempty"`
	Permiso          *Permiso  `gorm:"foreignKey:PermisoID;references:PermisoID" json:"permiso,omitempty"`
}
