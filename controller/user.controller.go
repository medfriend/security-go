package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"security-go/entity"
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

// CreateUser @Summary      Crear un nuevo usuario
// @Description  Este endpoint permite crear un nuevo usuario en el sistema
// @Security      BearerAuth
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Param        user  body      entity.User       true  "Información del usuario"
// @Success      201   {object}  entity.User
// @Router       /user [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
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
	util.HandlerCreatedSuccess(c, user)
}

// GetUsers obtiene todos los usuarios
// @Summary      Obtener todos los usuarios
// @Security      BearerAuth
// @Description  Este endpoint devuelve todos los usuarios del sistema.
// @Tags         usuarios
// @Produce      json
// @Success      200 {array}  entity.User   "Lista de usuarios"
// @Router       /user/all [get]
func (ctrl *UserController) GetUsers(c *gin.Context) {
	users, err := ctrl.userService.GetUsers()
	util.HandlerFoundSuccess(c, err, "usuarios")
	util.HandlerCreatedSuccess(c, users)
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
// @Router       /user [put]
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.userService.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser elimina un usuario por su ID
// @Summary      Eliminar un usuario por ID
// @Security      BearerAuth
// @Description  Este endpoint permite eliminar un usuario específico dado su ID.
// @Tags         usuarios
// @Param        id  path      uint  true  "ID del usuario"
// @Success      204 "Usuario eliminado con éxito"
// @Router       /user/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, _ := util.StringToUint(c.Param("id"))
	if err := ctrl.userService.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
