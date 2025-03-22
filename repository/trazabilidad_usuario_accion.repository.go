package repository

import (
	"github.com/medfriend/shared-commons-go/util/repository"
	"gorm.io/gorm"
	"security-go/entity"
)

type TrazabilidadUsuarioAccionRepository interface {
	Find(usuarioId uint) ([]entity.TrazabilidadUsuarioAccion, error)
}

type trazabilidadUsuarioAccionRepositoryImpl struct {
	Base repository.BaseRepository[entity.TrazabilidadUsuarioAccion]
}

func (t trazabilidadUsuarioAccionRepositoryImpl) Find(usuarioId uint) ([]entity.TrazabilidadUsuarioAccion, error) {
	var traza []entity.TrazabilidadUsuarioAccion
	result := t.Base.DB.Where("usuario_id = ?", usuarioId).
		Order("fecha_ejecucion DESC").
		Limit(5).
		Find(&traza)

	if result.Error != nil {
		return nil, result.Error
	}

	return traza, nil
}

func NewTrazabilidadUsuarioAccionRepository(db *gorm.DB) TrazabilidadUsuarioAccionRepository {
	return &trazabilidadUsuarioAccionRepositoryImpl{
		Base: repository.BaseRepository[entity.TrazabilidadUsuarioAccion]{DB: db},
	}
}
