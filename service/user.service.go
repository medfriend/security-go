package service

import (
	"security-go/entity"
	"security-go/repository"
	"security-go/util"
)

type UserService interface {
	CreateUser(user *entity.User) error
	GetUserById(id uint) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id uint) error
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

func (s *userServiceImpl) CreateUser(user *entity.User) error {
	hashedPassword, err := util.HashPassword(user.Clave)

	if err != nil {
		return err
	}

	user.Clave = hashedPassword

	return s.userRepo.Save(user)
}

func (s *userServiceImpl) GetUserById(id uint) (*entity.User, error) {
	return s.userRepo.FindById(id)
}

func (s *userServiceImpl) UpdateUser(user *entity.User) error {
	return s.userRepo.Update(user)
}

func (s *userServiceImpl) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
