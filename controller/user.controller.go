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
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Param        user  body      entity.User       true  "Información del usuario"
// @Success      201   {object}  entity.User
// @Router       /medfri-security/user [post]
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
// @Description  Este endpoint devuelve la información de un usuario específico dado su ID.
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Param        id  path      uint  true  "ID del usuario"
// @Success      200 {object}  entity.User   "Usuario encontrado"
// @Router       /medfri-security/user/{id} [get]
func (ctrl *UserController) GetUserById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))
	user, err := ctrl.userService.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser    actualiza la información de un usuario existente
// @Summary      Actualizar un usuario
// @Description  Este endpoint permite actualizar la información de un usuario existente.
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Param        user  body      entity.User  true  "Información del usuario"
// @Success      200   {object}  entity.User   "Usuario actualizado"
// @Router       /medfri-security/user [put]
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
// @Description  Este endpoint permite eliminar un usuario específico dado su ID.
// @Tags         usuarios
// @Param        id  path      uint  true  "ID del usuario"
// @Success      204 "Usuario eliminado con éxito"
// @Router       /medfri-securiry/user/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, _ := util.StringToUint(c.Param("id"))
	if err := ctrl.userService.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
