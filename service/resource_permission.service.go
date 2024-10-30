package service

import (
	"security-go/entity"
	"security-go/repository"
)

type ResourcePermissionService interface {
	CreateResourcePermission(resourcePermission *entity.ResourcePermission) error
	FindById(id uint) (*entity.ResourcePermission, error)
	UpdateResourcePermission(resourcePermission *entity.ResourcePermission) error
	DeleteResourcePermission(id uint) error
}

type ResourcePermissionServiceImpl struct {
	resourcePermissionRepository repository.ResourcePermissionRepository
}

func NewResourcePermissionService(resourcePermissionRepository repository.ResourcePermissionRepository) ResourcePermissionService {
	return &ResourcePermissionServiceImpl{
		resourcePermissionRepository: resourcePermissionRepository,
	}
}

func (r ResourcePermissionServiceImpl) CreateResourcePermission(resourcePermission *entity.ResourcePermission) error {
	return r.resourcePermissionRepository.Save(resourcePermission)
}

func (r ResourcePermissionServiceImpl) FindById(id uint) (*entity.ResourcePermission, error) {
	return r.resourcePermissionRepository.FindById(id)
}

func (r ResourcePermissionServiceImpl) UpdateResourcePermission(resourcePermission *entity.ResourcePermission) error {
	return r.resourcePermissionRepository.Update(resourcePermission)
}

func (r ResourcePermissionServiceImpl) DeleteResourcePermission(id uint) error {
	return r.resourcePermissionRepository.Delete(id)
}