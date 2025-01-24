package service

import (
	"security-go/entity"
	"security-go/repository"
)

type UserRolService interface {
	CreateUserRol(userRol *entity.UserRol) error
	DeleteUserRol(id uint) error
	FindRolesByUserID(userId uint) ([]uint, error)
	CheckUserRole(userId uint) (int64, error)
}

type UserRolServiceImpl struct {
	UserRolRepo repository.UserRolRepository
}

func NewUserRolService(UserRolRepo repository.UserRolRepository) UserRolService {
	return &UserRolServiceImpl{
		UserRolRepo: UserRolRepo,
	}
}

func (s *UserRolServiceImpl) CheckUserRole(userId uint) (int64, error) {
	return s.UserRolRepo.CheckUserRole(userId)
}

func (s *UserRolServiceImpl) FindRolesByUserID(userId uint) ([]uint, error) {
	return s.UserRolRepo.FindRolesByUserID(userId)
}

func (s *UserRolServiceImpl) CreateUserRol(UserRol *entity.UserRol) error {
	return s.UserRolRepo.Save(UserRol)
}

func (s *UserRolServiceImpl) DeleteUserRol(id uint) error {
	return s.UserRolRepo.Delete(id)
}
