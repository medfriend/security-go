package service

import (
	"fmt"
	"security-go/dto"
	"security-go/response"
	"security-go/util"
)

type AuthService interface {
	Auth(auth *dto.AuthDTO) (token *string, userId uint, err error)
}

type AuthServiceImpl struct {
	userRolService            UserRolService
	userService               UserService
	rolResourceService        RoleResourceService
	menuService               MenuService
	resourcePermissionService ResourcePermissionService
	EntityService             EntityService
}

func NewAuthService(
	userRolService UserRolService,
	userService UserService,
	rolResourceService RoleResourceService,
	menuService MenuService,
	resourcePermissionService ResourcePermissionService,
	entityService EntityService) AuthService {

	return &AuthServiceImpl{
		userRolService:            userRolService,
		userService:               userService,
		rolResourceService:        rolResourceService,
		menuService:               menuService,
		resourcePermissionService: resourcePermissionService,
		EntityService:             entityService,
	}
}

func (s *AuthServiceImpl) Auth(auth *dto.AuthDTO) (token *string, userId uint, err error) {

	user, _ := s.userService.FindByUsuario(auth.Usuario)

	if user == nil {
		return nil, 0, fmt.Errorf("usuario no encontrado")
	}

	passwordCheck := util.CheckPassword(user.Clave, auth.Password)

	if !passwordCheck {
		return nil, 0, fmt.Errorf("contraseña invalida")
	}

	entity, _ := s.userRolService.CheckUserRole(user.UsuarioID)

	roles, _ := s.userRolService.FindRolesByUserID(user.UsuarioID)

	resource, _ := s.rolResourceService.FindResourceByRoleIds(roles)

	permissions, _ := s.resourcePermissionService.FindPermissionByResourceAndRole(resource, roles)

	menus, _ := s.menuService.FindMenuByResourceAndEntity(resource, uint(entity), permissions)

	entityFound, _ := s.EntityService.GetEntityById(uint(entity))

	authResponse := response.AuthResponse{
		Menus:         *menus,
		User:          *user,
		EntidadId:     uint(entity),
		NombreEntidad: entityFound.RazonSocial,
	}

	fmt.Println(authResponse)

	/*rabbitMQ := util.GetInstance()

	userJson, err := json.Marshal(&user)

	rabbitMQ.SendMessage(
		"trazabilidad-usuario-login",
		string(userJson))*/

	jwt, _ := util.GenerateJWT(authResponse)

	return &jwt, user.UsuarioID, nil
}
