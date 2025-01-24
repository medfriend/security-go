package entity

func (Rol) TableName() string {
	return "rol"
}

type Rol struct {
	RolID       uint   `gorm:"primaryKey;autoIncrement;column:rol_id"`
	Nombre      string `gorm:"column:nombre" json:"nombre"`
	Descripcion string `gorm:"column:descripcion" json:"descripcion"`
	EntidadID   uint   `gorm:"column:entidad_id" json:"entidad_id"`
	Activo      bool   `gorm:"column:activo;default:true" json:"activo"`
}
