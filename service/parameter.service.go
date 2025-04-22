package service

import (
	"security-go/dto"
	"security-go/repository"
)

type ParameterService interface {
	FindParameterByCode(code string) (dto.FindParameterByCodeDto, error)
}

type ParameterServiceImpl struct {
	ParameterRepo repository.ParameterRepository
}

func (p ParameterServiceImpl) FindParameterByCode(code string) (dto.FindParameterByCodeDto, error) {
	return p.ParameterRepo.FindParameterByCode(code)
}

func NewParameterService(parameterRepo repository.ParameterRepository) ParameterService {
	return &ParameterServiceImpl{
		ParameterRepo: parameterRepo,
	}
}
