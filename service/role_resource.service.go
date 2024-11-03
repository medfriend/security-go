package service

import (
	"security-go/entity"
	"security-go/repository"
)

type RoleResourceService interface {
	CreateRoleResource(roleResource *entity.RoleResource) error
	FindById(id uint) (*entity.RoleResource, error)
	FindResourceByRoleIds(roleIds []uint) ([]uint, error)
	UpdateRoleResource(roleResource *entity.RoleResource) error
	DeleteRoleResource(id uint) error
}

type RoleResourceServiceImpl struct {
	roleResourceRepository repository.RoleResourceRepository
}

func NewRoleResourceService(roleResourceRepository repository.RoleResourceRepository) RoleResourceService {
	return &RoleResourceServiceImpl{
		roleResourceRepository: roleResourceRepository,
	}
}

func (r RoleResourceServiceImpl) CreateRoleResource(roleResource *entity.RoleResource) error {
	return r.roleResourceRepository.Save(roleResource)
}

func (r RoleResourceServiceImpl) FindById(id uint) (*entity.RoleResource, error) {
	return r.roleResourceRepository.FindById(id)
}

func (r RoleResourceServiceImpl) UpdateRoleResource(roleResource *entity.RoleResource) error {
	return r.roleResourceRepository.Update(roleResource)
}

func (r RoleResourceServiceImpl) DeleteRoleResource(id uint) error {
	return r.roleResourceRepository.Delete(id)
}

func (r RoleResourceServiceImpl) FindResourceByRoleIds(roleIds []uint) ([]uint, error) {
	return r.roleResourceRepository.FindResourceByRoleIds(roleIds)
}
