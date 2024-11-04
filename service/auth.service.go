package service

import (
	"fmt"
	"security-go/dto"
	"security-go/response"
	"security-go/util"
)

type AuthService interface {
	Auth(auth *dto.AuthDTO) (token *string, err error)
}

type AuthServiceImpl struct {
	userRolService            UserRolService
	userService               UserService
	rolResourceService        RoleResourceService
	menuService               MenuService
	resourcePermissionService ResourcePermissionService
}

func NewAuthService(
	userRolService UserRolService,
	userService UserService,
	rolResourceService RoleResourceService,
	menuService MenuService,
	resourcePermissionService ResourcePermissionService) AuthService {

	return &AuthServiceImpl{
		userRolService:            userRolService,
		userService:               userService,
		rolResourceService:        rolResourceService,
		menuService:               menuService,
		resourcePermissionService: resourcePermissionService,
	}
}

func (s *AuthServiceImpl) Auth(auth *dto.AuthDTO) (token *string, err error) {

	user, _ := s.userService.FindByUsuario(auth.Usuario)

	if user == nil {
		return nil, fmt.Errorf("usuario no encontrado")
	}

	passwordCheck := util.CheckPassword(user.Clave, auth.Password)

	if !passwordCheck {
		return nil, fmt.Errorf("contraseña invalida")
	}

	entity, _ := s.userRolService.CheckUserRole(user.UsuarioID)

	roles, _ := s.userRolService.FindRolesByUserID(user.UsuarioID)
	resource, _ := s.rolResourceService.FindResourceByRoleIds(roles)

	permissions, _ := s.resourcePermissionService.FindPermissionByResourceAndRole(resource, roles)

	menus, _ := s.menuService.FindMenuByResourceAndEntity(resource, uint(entity), permissions)

	authResponse := response.AuthResponse{
		Menus: *menus,
		User:  *user,
	}

	jwt, _ := util.GenerateJWT(authResponse)

	return &jwt, nil
}
