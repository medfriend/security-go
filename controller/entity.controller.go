package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
// @Description Este endpoint permite crear una nueva entidad en el sistema.
// @Tags entidades
// @Accept json
// @Produce json
// @Param entity body entity.Entity true "Información de la entidad"
// @Success 201 {object} entity.Entity "Entidad creada con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /medfri-security/entity [post]
func (ctrl *EntityController) CreateEntity(c *gin.Context) {
	var Entity entity.Entity
	if err := c.ShouldBindJSON(&Entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.EntityService.CreateEntity(&Entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, Entity)
}

// GetEntityById obtiene una entidad por su ID
// @Summary Obtener una entidad por ID
// @Description Este endpoint permite obtener la información de una entidad específica usando su ID.
// @Tags entidades
// @Accept json
// @Produce json
// @Param id path uint true "ID de la entidad"
// @Success 200 {object} entity.Entity "Entidad encontrada"
// @Failure 404 {object} map[string]string "Entidad no encontrada"
// @Router /medfri-security/entity/{id} [get]
func (ctrl *EntityController) GetEntityById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))
	Entity, err := ctrl.EntityService.GetEntityById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
		return
	}
	c.JSON(http.StatusOK, Entity)
}

// UpdateEntity actualiza una entidad existente
// @Summary Actualizar una entidad
// @Description Este endpoint permite actualizar la información de una entidad existente.
// @Tags entidades
// @Accept json
// @Produce json
// @Param entity body entity.Entity true "Información de la entidad actualizada"
// @Success 200 {object} entity.Entity "Entidad actualizada con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /medfri-security/entity [put]
func (ctrl *EntityController) UpdateEntity(c *gin.Context) {
	var Entity entity.Entity
	if err := c.ShouldBindJSON(&Entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.EntityService.UpdateEntity(&Entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Entity)
}

// DeleteEntity elimina una entidad por su ID
// @Summary Eliminar una entidad
// @Description Este endpoint permite eliminar una entidad específica usando su ID.
// @Tags entidades
// @Accept json
// @Produce json
// @Param id path uint true "ID de la entidad"
// @Success 204 "Entidad eliminada con éxito"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /medfri-security/entity/{id} [delete]
func (ctrl *EntityController) DeleteEntity(c *gin.Context) {
	id, _ := util.StringToUint(c.Param("id"))
	if err := ctrl.EntityService.DeleteEntity(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
