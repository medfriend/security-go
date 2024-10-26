package service

import (
	"security-go/entity"
	"security-go/repository"
)

type PermisoService interface {
	CreatePermiso(permiso *entity.Permiso) error
	GetPermisoById(id uint) (*entity.Permiso, error)
	UpdatePermiso(permiso *entity.Permiso) error
	DeletePermiso(id uint) error
}

type permisoServiceImpl struct {
	permisoRepo repository.PermisoRepository
}

func NewPermisoService(permisoRepo repository.PermisoRepository) PermisoService {
	return &permisoServiceImpl{
		permisoRepo: permisoRepo,
	}
}

func (s *permisoServiceImpl) CreatePermiso(permiso *entity.Permiso) error {
	return s.permisoRepo.Save(permiso)
}

func (s *permisoServiceImpl) GetPermisoById(id uint) (*entity.Permiso, error) {
	return s.permisoRepo.FindById(id)
}

func (s *permisoServiceImpl) UpdatePermiso(permiso *entity.Permiso) error {
	return s.permisoRepo.Update(permiso)
}

func (s *permisoServiceImpl) DeletePermiso(id uint) error {
	return s.permisoRepo.Delete(id)
}
