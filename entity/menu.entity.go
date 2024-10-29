package entity

func (Menu) TableName() string {
	return "menu"
}

type Menu struct {
	MenuID      uint   `gorm:"primaryKey;autoIncrement;column:menu_id"`
	Nombre      string `gorm:"column:nombre" json:"nombre"`
	MenuPadreID *uint  `gorm:"column:menu_padre_id" json:"menu_padre_id,omitempty"`
	RecursoID   *uint  `gorm:"column:recurso_id" json:"recurso_id,omitempty"`
	EntidadID   *uint  `gorm:"column:entidad_id" json:"entidad_id,omitempty"`
	Descripcion string `gorm:"column:descripcion" json:"descripcion"`

	MenuPadre *Menu     `gorm:"foreignKey:MenuPadreID;references:MenuID" json:"menu_padre,omitempty"`
	Recurso   *Resource `gorm:"foreignKey:RecursoID" json:"recurso,omitempty"`
	Entidad   *Entity   `gorm:"foreignKey:EntidadID" json:"entidad,omitempty"`
}
