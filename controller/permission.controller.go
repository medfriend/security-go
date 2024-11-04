package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"security-go/entity"
	"security-go/service"
	"security-go/util"
)

type PermisoController struct {
	permisoService service.PermisoService
}

func NewPermisoController(permisoService service.PermisoService) *PermisoController {
	return &PermisoController{
		permisoService: permisoService,
	}
}

// CreatePermiso crea una nueva entidad
// @Summary Crear un permiso
// @Description Este endpoint permite crear un nuevo permiso.
// @Tags permisos
// @Accept json
// @Produce json
// @Param entity body entity.Permiso true "Información de la permisos"
// @Success 201 {object} entity.Permiso "permiso creada con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /medfri-security/permission [post]
func (ctrl *PermisoController) CreatePermiso(c *gin.Context) {
	var permiso entity.Permiso
	util.HandlerBindJson(c, &permiso)
	util.HandlerInternalError(c, ctrl.permisoService.CreatePermiso(&permiso))
	util.HandlerCreatedSuccess(c, permiso)
}

// GetPermisoById obtiene un permiso por su ID
// @Summary Obtener una permiso por ID
// @Description Este endpoint permite obtener la información de una permiso específica usando su ID.
// @Tags permisos
// @Accept json
// @Produce json
// @Param id path uint true "ID de la permiso"
// @Success 200 {object} entity.Permiso "Permiso encontrada"
// @Failure 404 {object} map[string]string "Permiso no encontrada"
// @Router /medfri-security/permission/{id} [get]
func (ctrl *PermisoController) GetPermisoById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))
	permiso, err := ctrl.permisoService.GetPermisoById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Permiso not found"})
		return
	}
	c.JSON(http.StatusOK, permiso)
}

// UpdatePermiso actualiza un permiso existente
// @Summary Actualizar un permiso
// @Description Este endpoint permite actualizar la información de una permiso existente.
// @Tags permisos
// @Accept json
// @Produce json
// @Param entity body entity.Permiso true "Información de la permisos actualizada"
// @Success 200 {object} entity.Permiso "permiso actualizada con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /medfri-security/permission [put]
func (ctrl *PermisoController) UpdatePermiso(c *gin.Context) {
	var permiso entity.Permiso
	if err := c.ShouldBindJSON(&permiso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.permisoService.UpdatePermiso(&permiso); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, permiso)
}

// DeletePermiso elimina un permiso por su ID
// @Summary Eliminar un permiso
// @Description Este endpoint permite eliminar un permiso específica usando su ID.
// @Tags permisos
// @Accept json
// @Produce json
// @Param id path uint true "ID de la permisos"
// @Success 204 "Permisos eliminada con éxito"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /medfri-security/permission/{id} [delete]
func (ctrl *PermisoController) DeletePermiso(c *gin.Context) {
	id, _ := util.StringToUint(c.Param("id"))
	if err := ctrl.permisoService.DeletePermiso(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
