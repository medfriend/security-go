package service

import (
	"fmt"
	"security-go/dto"
)

type AuthService interface {
	Auth(auth *dto.AuthDTO) error
}

type AuthServiceImpl struct {
	userRolService UserRolService
	userService    UserService
}

func NewAuthService(userRolService UserRolService, userService UserService) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRolService: userRolService,
		userService:    userService,
	}
}

func (s *AuthServiceImpl) Auth(auth *dto.AuthDTO) error {
	fmt.Println(auth)
	user, err := s.userService.FindByUsuario(auth.Usuario)

	if err != nil {
		return err
	}

	roles, err := s.userRolService.FindRolesByUserID(user.UsuarioID)

	if err != nil {
		return err
	}

	fmt.Println(roles)

	return nil
}
