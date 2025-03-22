package service

import (
	"security-go/entity"
	"security-go/repository"
)

type TrazabilidadUsuarioAccionService interface {
	Find(usuarioId uint) ([]entity.TrazabilidadUsuarioAccion, error)
}

type TrazabilidadUsuarioAccionServiceImpl struct {
	trazaRepo repository.TrazabilidadUsuarioAccionRepository
}

func (t TrazabilidadUsuarioAccionServiceImpl) Find(usuarioId uint) ([]entity.TrazabilidadUsuarioAccion, error) {
	return t.trazaRepo.Find(usuarioId)
}

func NewTrazabilidadUsuarioAccionService(trazaRepo repository.TrazabilidadUsuarioAccionRepository) TrazabilidadUsuarioAccionService {
	return &TrazabilidadUsuarioAccionServiceImpl{
		trazaRepo: trazaRepo,
	}
}
