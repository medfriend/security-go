package controller

import (
	"github.com/gin-gonic/gin"
	"security-go/entity"
	"security-go/service"
	"security-go/util"
)

type EntityController struct {
	EntityService service.EntityService
}

func NewEntityController(EntityService service.EntityService) *EntityController {
	return &EntityController{
		EntityService: EntityService,
	}
}

// CreateEntity crea una nueva entidad
// @Summary Crear una entidad
// @Security      BearerAuth
// @Description Este endpoint permite crear una nueva entidad en el sistema.
// @Tags entidades
// @Accept json
// @Produce json
// @Param entity body entity.Entity true "Información de la entidad"
// @Success 201 {object} entity.Entity "Entidad creada con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /entity [post]
func (ctrl *EntityController) CreateEntity(c *gin.Context) {
	var Entity entity.Entity

	util.HandlerBindJson(c, &Entity)
	util.HandlerInternalError(c, ctrl.EntityService.CreateEntity(&Entity))
	util.HandlerCreatedSuccess(c, Entity)
}

// GetEntityById obtiene una entidad por su ID
// @Summary Obtener una entidad por ID
// @Security      BearerAuth
// @Description Este endpoint permite obtener la información de una entidad específica usando su ID.
// @Tags entidades
// @Accept json
// @Produce json
// @Param id path uint true "ID de la entidad"
// @Success 200 {object} entity.Entity "Entidad encontrada"
// @Failure 404 {object} map[string]string "Entidad no encontrada"
// @Router /entity/{id} [get]
func (ctrl *EntityController) GetEntityById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))
	Entity, err := ctrl.EntityService.GetEntityById(id)
	util.HandlerFoundSuccess(c, err, "entidad")
	util.HandlerCreatedSuccess(c, Entity)
}

// GetAllEntities obtiene todas las entidades registradas
// @Summary Obtener todas las entidades
// @Security      BearerAuth
// @Description Este endpoint obtener todas las entidades
// @Tags entidades
// @Accept json
// @Produce json
// @Success 200 {array} entity.Entity "Entidades encontrada"
// @Failure 404 {object} map[string]string "Entidades no encontrada"
// @Router /entity/all [get]
func (ctrl *EntityController) GetAllEntities(c *gin.Context) {
	Entity, err := ctrl.EntityService.GetAllEntities()
	util.HandlerFoundSuccess(c, err, "entidad")
	util.HandlerCreatedSuccess(c, Entity)
}

// UpdateEntity actualiza una entidad existente
// @Summary Actualizar una entidad
// @Security      BearerAuth
// @Description Este endpoint permite actualizar la información de una entidad existente.
// @Tags entidades
// @Accept json
// @Produce json
// @Param entity body entity.Entity true "Información de la entidad actualizada"
// @Success 200 {object} entity.Entity "Entidad actualizada con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /entity [put]
func (ctrl *EntityController) UpdateEntity(c *gin.Context) {
	var Entity entity.Entity
	util.HandlerBindJson(c, &Entity)
	util.HandlerInternalError(c, ctrl.EntityService.UpdateEntity(&Entity))
	util.HandlerCreatedSuccess(c, Entity)
}

// DeleteEntity elimina una entidad por su ID
// @Summary Eliminar una entidad
// @Security      BearerAuth
// @Description Este endpoint permite eliminar una entidad específica usando su ID.
// @Tags entidades
// @Accept json
// @Produce json
// @Param id path uint true "ID de la entidad"
// @Success 204 "Entidad eliminada con éxito"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /entity/{id} [delete]
func (ctrl *EntityController) DeleteEntity(c *gin.Context) {
	id, _ := util.StringToUint(c.Param("id"))
	util.HandlerInternalError(c, ctrl.EntityService.DeleteEntity(id))
	util.HandlerNotContent(c, nil)
}
