package entity

func (Resource) TableName() string {
	return "recurso"
}

type Resource struct {
	ResourceID  uint   `gorm:"primaryKey;autoIncrement;column:recurso_id" json:"recurso_id"`
	Descripcion string `gorm:"size:100;not null;column:descripcion" json:"descripcion"`
	Acceso      string `gorm:"size:100;column:acceso" json:"acceso"`
}
