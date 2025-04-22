package service

import (
	"fmt"
	"security-go/dto"
	"security-go/entity"
	"security-go/repository"
)

type RolService interface {
	CreateRol(Rol *entity.Rol) error
	FindById(id uint) (*entity.Rol, error)
	Find(paginacion dto.PaginationDTO) (dto.PaginatedResponse, error)
	UpdateRol(Rol *entity.Rol) error
	DeleteRol(id uint) error
}

type RolServiceImpl struct {
	RolRepository repository.RolRepository
}

func (m RolServiceImpl) Find(paginacion dto.PaginationDTO) (dto.PaginatedResponse, error) {
	return m.RolRepository.Find(paginacion)
}

func NewRolService(RolRepository repository.RolRepository) RolService {
	return &RolServiceImpl{
		RolRepository: RolRepository,
	}
}

func (m RolServiceImpl) CreateRol(Rol *entity.Rol) error {
	fmt.Println(Rol)
	return m.RolRepository.Save(Rol)
}

func (m RolServiceImpl) FindById(id uint) (*entity.Rol, error) {
	return m.RolRepository.FindById(id)
}

func (m RolServiceImpl) UpdateRol(Rol *entity.Rol) error {
	return m.RolRepository.Update(Rol)
}

func (m RolServiceImpl) DeleteRol(id uint) error {
	return m.RolRepository.Delete(id)
}
