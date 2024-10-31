package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"security-go/entity"
	"security-go/service"
	"security-go/util"
)

type UserRolController struct {
	UserRolService service.UserRolService
}

func NewUserRolController(UserRolService service.UserRolService) *UserRolController {
	return &UserRolController{
		UserRolService: UserRolService,
	}
}

// CreateUserRol crea un nuevo relacion usuario rol
// @Summary Crear un recurso
// @Description Este endpoint permite crear un nueva relacion usuarioRol en el sistema.
// @Tags re-usuario-rol
// @Accept json
// @Produce json
// @Param UserRol body entity.UserRol true "Información del recurso"
// @Success 201 {object} entity.UserRol "Recurso creado con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /medfri-security/user-rol/asignar [post]
func (ctrl *UserRolController) CreateUserRol(c *gin.Context) {
	var userRol entity.UserRol

	util.HandlerBindJson(c, &userRol)
	util.HandlerInternalError(c, ctrl.UserRolService.CreateUserRol(&userRol))

	c.JSON(http.StatusCreated, userRol)
}

// DeleteUserRol elimina una relacion usuarioRol por su ID
// @Summary Eliminar una relacion usuario rol
// @Description Este endpoint permite eliminar una relacion usuario rol específico usando su ID.
// @Tags re-usuario-rol
// @Accept json
// @Produce json
// @Param id path uint true "ID del usuario rol"
// @Success 204 "usuario rol eliminado con éxito"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /medfri-security/user-rol/desasignar/{id} [delete]
func (ctrl *UserRolController) DeleteUserRol(c *gin.Context) {
	var userRol entity.UserRol
	id, _ := util.StringToUint(c.Param("id"))

	util.HandlerBindJson(c, &userRol)
	util.HandlerInternalError(c, ctrl.UserRolService.DeleteUserRol(id))

	c.JSON(http.StatusNoContent, nil)
}
