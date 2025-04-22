package entity

type Parameter struct {
	ParametroID      int64  `gorm:"column:parametro_id;primaryKey;autoIncrement" json:"parametro_id"`
	Nombre           string `gorm:"column:nombre;type:varchar(255)" json:"nombre"`
	Descripcion      string `gorm:"column:descripcion;type:varchar(255)" json:"descripcion"`
	Codigo           string `gorm:"column:codigo;type:varchar(255)" json:"codigo"`
	ParametroPadreID *int64 `gorm:"column:parametro_padre_id" json:"parametro_padre_id"`
}

func (Parameter) TableName() string {
	return "parametro"
}
