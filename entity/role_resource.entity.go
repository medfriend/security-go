package entity

func (RoleResource) TableName() string {
	return "re-rol-recurso"
}

type RoleResource struct {
	RolRecursoID uint      `gorm:"primaryKey;autoIncrement;column:rol_recurso_id"`
	RolID        *uint     `gorm:"column:rol_id" json:"rol_id,omitempty"`
	RecursoID    *uint     `gorm:"column:recurso_id" json:"recurso_id,omitempty"`
	EntidadID    *uint     `gorm:"column:entidad_id" json:"entidad_id,omitempty"`
	Acceso       string    `gorm:"column:acceso" json:"acceso"`
	Rol          *Rol      `gorm:"foreignKey:RolID" json:"rol,omitempty"`
	Recurso      *Resource `gorm:"foreignKey:RecursoID" json:"recurso,omitempty"`
	Entidad      *Entity   `gorm:"foreignKey:EntidadID" json:"entidad,omitempty"`
}