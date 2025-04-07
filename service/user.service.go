package service

import (
	"fmt"
	"security-go/dto"
	"security-go/entity"
	"security-go/repository"
	"security-go/util"
)

type UserService interface {
	CreateUser(user *entity.User) error
	GetUserById(id uint) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	UpdateUser(user *dto.UpdateUserDTO) error
	DeleteUser(id uint) error
	FindByUsuario(usuario uint) (*entity.User, error)
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

	fmt.Println(user)

	return nil
}

func (s *userServiceImpl) GetUserById(id uint) (*entity.User, error) {
	return s.userRepo.FindById(id)
}

func (s *userServiceImpl) GetUsers() ([]entity.User, error) {
	return s.userRepo.Find()
}

func (s *userServiceImpl) UpdateUser(user *dto.UpdateUserDTO) error {
	return s.userRepo.Update(user)
}

func (s *userServiceImpl) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

func (s *userServiceImpl) FindByUsuario(usuario uint) (*entity.User, error) {
	return s.userRepo.FindByUsuario(usuario)
}
