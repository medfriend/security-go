package controller

import (
	"net/http"
	"security-go/entity"
	"security-go/service"
	"security-go/util"

	"github.com/gin-gonic/gin"
)

type RoleResourceController struct {
	RoleResource service.RoleResourceService
}

func NewRoleResourceController(roleResourceService service.RoleResourceService) *RoleResourceController {
	return &RoleResourceController{
		RoleResource: roleResourceService,
	}
}

// CreateRoleResource crea un nuevo rol recurso
// @Summary Crear un rol recurso
// @Security      BearerAuth
// @Description Este endpoint permite crear un nuevo rol recurso en el sistema.
// @Tags roles-recursos
// @Accept json
// @Produce json
// @Param resource body entity.RoleResource true "Información del rol recurso"
// @Success 201 {object} entity.RoleResource "Rol recurso creado con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /rol-recurso/asignar [post]
func (ctrl *RoleResourceController) CreateRoleResource(c *gin.Context) {
	var roleResource entity.RoleResource

	util.HandlerBindJson(c, &roleResource)
	util.HandlerInternalError(c, ctrl.RoleResource.CreateRoleResource(&roleResource))
	c.JSON(http.StatusCreated, roleResource)
}

// GetRoleResourceById   obtiene un rol recurso por su ID
// @Summary      Obtener un rol recurso por ID
// @Security      BearerAuth
// @Description  Este endpoint devuelve la información de un rol recurso específico dado su ID.
// @Tags         roles-recursos
// @Accept       json
// @Produce      json
// @Param        id  path      uint  true  "ID del rol recurso"
// @Success      200 {object}  entity.RoleResource   "rol recurso encontrado"
// @Router       /rol-recurso/{id} [get]
func (ctrl *RoleResourceController) GetRoleResourceById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))

	roleResource, err := ctrl.RoleResource.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "RoleResource not found"})
		return
	}
	c.JSON(http.StatusOK, roleResource)
}

// UpdateRoleResource actualiza un rol recurso
// @Summary Actualizar un rol recurso
// @Security      BearerAuth
// @Description Este endpoint permite actualizar la información de un rol recurso existente.
// @Tags roles-recursos
// @Accept json
// @Produce json
// @Param resource body entity.RoleResource true "Información del rol recurso actualizada"
// @Success 200 {object} entity.RoleResource "rol recurso actualizado con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /rol-recurso [put]
func (ctrl *RoleResourceController) UpdateRoleResource(c *gin.Context) {
	var roleResource entity.RoleResource
	if err := c.ShouldBindJSON(&roleResource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.RoleResource.UpdateRoleResource(&roleResource); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roleResource)
}

// DeleteRoleResource elimina un rol recurso
// @Summary Eliminar un rol recurso
// @Security      BearerAuth
// @Description Este endpoint permite eliminar un rol recurso existente en el sistema.
// @Tags roles-recursos
// @Accept json
// @Produce json
// @Param id path uint true "ID del rol recurso"
// @Success 200 {string} string "Rol recurso eliminado con éxito"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /rol-recurso/desasignar/{id} [delete]
func (ctrl *RoleResourceController) DeleteRoleResource(c *gin.Context) {
	id, _ := util.StringToUint(c.Param("id"))
	if err := ctrl.RoleResource.DeleteRoleResource(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "RoleResource deleted successfully")
}
