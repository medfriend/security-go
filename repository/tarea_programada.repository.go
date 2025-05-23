package repository

import (
	"gorm.io/gorm"
	"security-go/entity"
	"security-go/util"
)

type TareaProgramadaRepository interface {
	Find() (*[]entity.TareaProgramada, error)
}

type TareaProgramadaRepositoryImpl struct {
	Base util.BaseRepository[entity.TareaProgramada]
}

func (t TareaProgramadaRepositoryImpl) Find() (*[]entity.TareaProgramada, error) {
	var tareas []entity.TareaProgramada
	err := t.Base.DB.Where("activo = ?", true).Find(&tareas).Error

	return &tareas, err
}

func NewTareaProgramadaRepository(db *gorm.DB) TareaProgramadaRepository {
	return &TareaProgramadaRepositoryImpl{
		Base: util.BaseRepository[entity.TareaProgramada]{DB: db},
	}
}
