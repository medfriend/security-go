package entity

type Hospital struct {
	HospitalID  int64  `gorm:"primaryKey;autoIncrement" json:"hospital_id"`
	NIT         int64  `gorm:"not null" json:"nit"`
	NITAlterno  string `gorm:"type:varchar(254)" json:"nit_alterno"`
	RazonSocial string `gorm:"type:varchar(254);not null" json:"razon_social"`
	Email       string `gorm:"type:varchar(50)" json:"email"`
	Activo      bool   `gorm:"default:true" json:"activo"`
}
