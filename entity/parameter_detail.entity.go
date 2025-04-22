package entity

type ParameterDetail struct {
	ParametroDetalleID      int64  `gorm:"column:parametro_detalle_id;primaryKey;autoIncrement" json:"parametro_detalle_id"`
	ParametroID             int64  `gorm:"column:parametro_id" json:"parametro_id"`
	ValorCodigo             string `gorm:"column:valor_codigo;type:varchar(255)" json:"valor_codigo"`
	ValorDescripcion        string `gorm:"column:valor_descripcion;type:varchar(255)" json:"valor_descripcion"`
	Estado                  string `gorm:"column:estado;type:varchar(50)" json:"estado"`
	ParametroDetallePadreID *int64 `gorm:"column:parametro_detalle_padre_id" json:"parametro_detalle_padre_id"`
}

func (ParameterDetail) TableName() string {
	return "parametro_detalle"
}
