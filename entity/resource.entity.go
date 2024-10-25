package entity

import (
	"gorm.io/gorm"
)

type Resource struct {
	ResourceID  uint   gorm:"primaryKey;autoIncrement" json:"resource_id"
	description string gorm:"size:100;not null" json:"description"
	access      string gorm:"size:100" json:"access"

	gorm.Model
}