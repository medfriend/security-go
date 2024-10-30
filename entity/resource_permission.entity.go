package entity

func (ResourcePermission) TableName() string {
	return "re-recurso-permisos"
}

type ResourcePermission struct {
	RecursoPermisoID uint      `gorm:"primaryKey;autoIncrement;column:resource_permiso_id"`
	RecursoID        *uint     `gorm:"column:recurso_id" json:"recurso_id,omitempty"`
	PermisoID        *uint     `gorm:"column:permiso_id" json:"permiso_id,omitempty"`
	Recurso          *Resource `gorm:"foreignKey:RecursoID" json:"recurso,omitempty"`
	Permiso          *Permiso  `gorm:"foreignKey:PermisoID" json:"permiso,omitempty"`
}