package entity

func (Entity) TableName() string {
	return "entidad"
}

type Entity struct {
	EntityID    uint   `gorm:"primaryKey;autoIncrement;column:entidad_id" json:"entidad_id"`
	NIT         uint   `gorm:"not null;column:nit" json:"nit"`
	NITAlterno  string `gorm:"type:varchar(254);column:nit_alterno" json:"nit_alterno"`
	RazonSocial string `gorm:"type:varchar(254);not null;column:razon_social" json:"razon_social"`
	Email       string `gorm:"type:varchar(50);column:email" json:"email"`
	Activo      bool   `gorm:"default:true;column:activo" json:"activo"`
}
