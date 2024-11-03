package service

import (
	"fmt"
	"security-go/dto"
	"security-go/response"
)

type AuthService interface {
	Auth(auth *dto.AuthDTO) (menu *[]response.MenuResponse, err error)
}

type AuthServiceImpl struct {
	userRolService     UserRolService
	userService        UserService
	rolResourceService RoleResourceService
	menuService        MenuService
}

func NewAuthService(
	userRolService UserRolService,
	userService UserService,
	rolResourceService RoleResourceService,
	menuService MenuService) AuthService {

	return &AuthServiceImpl{
		userRolService:     userRolService,
		userService:        userService,
		rolResourceService: rolResourceService,
		menuService:        menuService,
	}
}

func (s *AuthServiceImpl) Auth(auth *dto.AuthDTO) (menu *[]response.MenuResponse, err error) {
	fmt.Println(auth)
	user, _ := s.userService.FindByUsuario(auth.Usuario)
	roles, _ := s.userRolService.FindRolesByUserID(user.UsuarioID)
	resource, _ := s.rolResourceService.FindResourceByRoleIds(roles)
	menus, _ := s.menuService.FindMenuByResourceAndEntity(resource, 1)

	fmt.Println(menus)

	return menus, nil
}
