package service

import (
	"fmt"
	"security-go/dto"
)

type AuthService interface {
	Auth(auth *dto.AuthDTO) error
}

type AuthServiceImpl struct {
	userRolService     UserRolService
	userService        UserService
	rolResourceService RoleResourceService
}

func NewAuthService(
	userRolService UserRolService,
	userService UserService,
	rolResourceService RoleResourceService) AuthService {

	return &AuthServiceImpl{
		userRolService:     userRolService,
		userService:        userService,
		rolResourceService: rolResourceService,
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
