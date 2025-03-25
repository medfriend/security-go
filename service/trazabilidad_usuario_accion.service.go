package service

import (
	"security-go/dto"
	"security-go/mapper"
	"security-go/repository"
)

type TrazabilidadUsuarioAccionService interface {
	Find(usuarioId uint) ([]dto.TrazaDTO, error)
}

type TrazabilidadUsuarioAccionServiceImpl struct {
	trazaRepo repository.TrazabilidadUsuarioAccionRepository
}

func (t TrazabilidadUsuarioAccionServiceImpl) Find(usuarioId uint) ([]dto.TrazaDTO, error) {
	trazas, err := t.trazaRepo.Find(usuarioId)
	return mapper.MapTrazabilidadToTrazaDTOArray(trazas), err
}

func NewTrazabilidadUsuarioAccionService(trazaRepo repository.TrazabilidadUsuarioAccionRepository) TrazabilidadUsuarioAccionService {
	return &TrazabilidadUsuarioAccionServiceImpl{
		trazaRepo: trazaRepo,
	}
}
