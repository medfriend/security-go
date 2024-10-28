package entity

func (Menu) TableName() string {
	return "menu"
}

type Menu struct {
	ItemID      uint   `gorm:"primaryKey;column:item_id" json:"item_id"`
	Nombre      string `gorm:"column:nombre" json:"nombre"`
	ItemPadreId *uint  `gorm:"column:item_padre_id" json:"item_padre_id,omitempty"`
	RecursoID   uint   `gorm:"column:recurso_id" json:"recurso_id"`
	EntidadID   uint   `gorm:"column:entidad_id" json:"entidad_id"`
	Descripcion string `gorm:"column:descripcion" json:"descripcion"`

	ItemPadre *Menu    `gorm:"foreignKey:ItemPadreID;references:ItemID" json:"item_padre,omitempty"`
	Recurso   Resource `gorm:"foreignKey:RecursoID" json:"recurso"`
	Entidad   Entity   `gorm:"foreignKey:EntidadID" json:"entidad"`
}
