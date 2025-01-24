package service

import (
	"fmt"
	"security-go/entity"
	"security-go/repository"
)

type RolService interface {
	CreateRol(Rol *entity.Rol) error
	FindById(id uint) (*entity.Rol, error)
	Find() ([]entity.Rol, error)
	UpdateRol(Rol *entity.Rol) error
	DeleteRol(id uint) error
}

type RolServiceImpl struct {
	RolRepository repository.RolRepository
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

func (m RolServiceImpl) Find() ([]entity.Rol, error) { return m.RolRepository.Find() }

func (m RolServiceImpl) FindById(id uint) (*entity.Rol, error) {
	return m.RolRepository.FindById(id)
}

func (m RolServiceImpl) UpdateRol(Rol *entity.Rol) error {
	return m.RolRepository.Update(Rol)
}

func (m RolServiceImpl) DeleteRol(id uint) error {
	return m.RolRepository.Delete(id)
}
