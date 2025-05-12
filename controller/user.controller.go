package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/medfriend/shared-commons-go/util/controller"
	"log"
	"security-go/dto"
	"security-go/mapper"
	"security-go/service"
	"security-go/util"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// FilterUser obtener los usuarios por medio de la query
// @Summary filtro de user
// @Security      BearerAuth
// @Description obtener menus por medio de la query
// @Param query path string true "campo de concidencia"
// @Tags menus
// @Accept json
// @Produce json
// @Success 200 {object} []entity.User "listado de usuarios"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /menu/filterUser/{query} [get]
func (ctrl *UserController) FilterUser(c *gin.Context) {
	query := c.Param("query")
	usuarios, err := ctrl.userService.Query(query)
	util.HandlerFoundSuccess(c, err, "usuarios concidentes")
	util.HandlerCreatedSuccess(c, usuarios, 0)
}

// CreateUser @Summary      Crear un nuevo usuario
// @Description  Este endpoint permite crear un nuevo usuario en el sistema
// @Security      BearerAuth
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Param        user  body      entity.User       true  "Información del usuario"
// @Success      201   {object}  entity.User
// @Router       /user/createuser [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var userDTO dto.UserDTO

	util.HandlerBindJson(c, &userDTO)
	user, _ := mapper.UserDTOToUser(userDTO)

	util.HandlerInternalError(c, ctrl.userService.CreateUser(user))
	util.HandlerCreatedSuccess(c, user, user.UsuarioID)
}

// GetUserById obtiene un usuario por su ID
// @Summary      Obtener un usuario por ID
// @Security      BearerAuth
// @Description  Este endpoint devuelve la información de un usuario específico dado su ID.
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Param        id  path      uint  true  "ID del usuario"
// @Success      200 {object}  entity.User   "Usuario encontrado"
// @Router       /user/byId/{id} [get]
func (ctrl *UserController) GetUserById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))
	user, err := ctrl.userService.GetUserById(id)
	util.HandlerFoundSuccess(c, err, "usuario")
	util.HandlerCreatedSuccess(c, user, user.UsuarioID)
}

// GetUsers obtiene todos los usuarios
// @Summary      Obtener todos los usuarios
// @Security      BearerAuth
// @Description  Este endpoint devuelve todos los usuarios del sistema.
// @Tags         usuarios
// @Produce      json
// @Success      200 {array}  entity.User   "Lista de usuarios"
// @Router       /user/all [post]
func (ctrl *UserController) GetUsers(c *gin.Context) {
	var paginacion dto.PaginationDTO
	controller.HandlerBindJson(c, &paginacion)

	authHeader := c.GetHeader("Authorization")
	log.Println("log de prueba")
	fmt.Println(authHeader)

	users, err := ctrl.userService.GetUsers(paginacion)
	util.HandlerFoundSuccess(c, err, "usuarios")
	util.HandlerCreatedSuccess(c, users, 0)
}

// UpdateUser    actualiza la información de un usuario existente
// @Summary      Actualizar un usuario
// @Security      BearerAuth
// @Description  Este endpoint permite actualizar la información de un usuario existente.
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Param        user  body      entity.User  true  "Información del usuario"
// @Success      200   {object}  entity.User   "Usuario actualizado"
// @Router       /user/update [put]
func (ctrl *UserController) UpdateUser(c *gin.Context) {

	var userDTO dto.UpdateUserDTO
	controller.HandlerBindJson(c, &userDTO)
	controller.HandlerInternalError(c, ctrl.userService.UpdateUser(&userDTO))
	util.HandlerCreatedSuccess(c, userDTO, uint(userDTO.Usuario_id))
}

// DeleteUser elimina un usuario por su ID
// @Summary      Eliminar un usuario por ID
// @Security      BearerAuth
// @Description  Este endpoint permite eliminar un usuario específico dado su ID.
// @Tags         usuarios
// @Param        id  path      uint  true  "ID del usuario"
// @Success      204 "Usuario eliminado con éxito"
// @Router       /user/deleteuser/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, _ := util.StringToUint(c.Param("id"))
	user, _ := ctrl.userService.GetUserById(id)
	controller.HandlerInternalError(c, ctrl.userService.DeleteUser(id))
	controller.HandlerDeleteSuccess(c, user, id)
}
