package controller

import (
	"github.com/gin-gonic/gin"
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
// @Security      BearerAuth
// @Description Este endpoint permite crear un nuevo permiso.
// @Tags permisos
// @Accept json
// @Produce json
// @Param entity body entity.Permiso true "Información de la permisos"
// @Success 201 {object} entity.Permiso "permiso creada con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /permission/createPermiso [post]
func (ctrl *PermisoController) CreatePermiso(c *gin.Context) {
	var permiso entity.Permiso
	util.HandlerBindJson(c, &permiso)
	util.HandlerInternalError(c, ctrl.permisoService.CreatePermiso(&permiso))
	util.HandlerCreatedSuccess(c, permiso, permiso.PermisoID)
}

// GetPermisoById obtiene un permiso por su ID
// @Summary Obtener una permiso por ID
// @Security      BearerAuth
// @Description Este endpoint permite obtener la información de una permiso específica usando su ID.
// @Tags permisos
// @Accept json
// @Produce json
// @Param id path uint true "ID de la permiso"
// @Success 200 {object} entity.Permiso "Permiso encontrada"
// @Failure 404 {object} map[string]string "Permiso no encontrada"
// @Router /permission/getPermisoById/{id} [get]
func (ctrl *PermisoController) GetPermisoById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))
	permiso, err := ctrl.permisoService.GetPermisoById(id)
	util.HandlerFoundSuccess(c, err, "permiso")
	util.HandlerCreatedSuccess(c, permiso, permiso.PermisoID)
}

// UpdatePermiso actualiza un permiso existente
// @Summary Actualizar un permiso
// @Security      BearerAuth
// @Description Este endpoint permite actualizar la información de una permiso existente.
// @Tags permisos
// @Accept json
// @Produce json
// @Param entity body entity.Permiso true "Información de la permisos actualizada"
// @Success 200 {object} entity.Permiso "permiso actualizada con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /permission/UpdatePermiso [put]
func (ctrl *PermisoController) UpdatePermiso(c *gin.Context) {
	var permiso entity.Permiso
	util.HandlerBindJson(c, &permiso)
	util.HandlerInternalError(c, ctrl.permisoService.UpdatePermiso(&permiso))
	util.HandlerCreatedSuccess(c, permiso, permiso.PermisoID)
}

// DeletePermiso elimina un permiso por su ID
// @Summary Eliminar un permiso
// @Security      BearerAuth
// @Description Este endpoint permite eliminar un permiso específica usando su ID.
// @Tags permisos
// @Accept json
// @Produce json
// @Param id path uint true "ID de la permisos"
// @Success 204 "Permisos eliminada con éxito"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /permission/deletePermiso/{id} [delete]
func (ctrl *PermisoController) DeletePermiso(c *gin.Context) {
	id, _ := util.StringToUint(c.Param("id"))
	util.HandlerInternalError(c, ctrl.permisoService.DeletePermiso(id))
	util.HandlerNotContent(c, nil)
}
