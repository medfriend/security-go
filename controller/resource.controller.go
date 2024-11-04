package controller

import (
	"net/http"
	"security-go/entity"
	"security-go/service"
	"security-go/util"

	"github.com/gin-gonic/gin"
)

type ResourceController struct {
	resourceService service.ResourceService
}

func NewResourceController(resourceService service.ResourceService) *ResourceController {
	return &ResourceController{
		resourceService: resourceService,
	}
}

// CreateResource crea un nuevo recurso
// @Summary Crear un recurso
// @Security      BearerAuth
// @Description Este endpoint permite crear un nuevo recurso en el sistema.
// @Tags recursos
// @Accept json
// @Produce json
// @Param resource body entity.Resource true "Información del recurso"
// @Success 201 {object} entity.Resource "Recurso creado con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /resources [post]
func (ctrl *ResourceController) CreateResource(c *gin.Context) {
	var resource entity.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.resourceService.CreateResource(&resource); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resource)
}

// GetResourceById obtiene un recurso por su ID
// @Summary Obtener un recurso por ID
// @Security      BearerAuth
// @Description Este endpoint permite obtener la información de un recurso específico usando su ID.
// @Tags recursos
// @Accept json
// @Produce json
// @Param id path uint true "ID del recurso"
// @Success 200 {object} entity.Resource "Recurso encontrado"
// @Failure 404 {object} map[string]string "Recurso no encontrado"
// @Router /resources/{id} [get]
func (ctrl *ResourceController) GetResourceById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))
	resource, err := ctrl.resourceService.GetResourceById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		return
	}
	c.JSON(http.StatusOK, resource)
}

// UpdateResource actualiza un recurso existente
// @Summary Actualizar un recurso
// @Security      BearerAuth
// @Description Este endpoint permite actualizar la información de un recurso existente.
// @Tags recursos
// @Accept json
// @Produce json
// @Param resource body entity.Resource true "Información del recurso actualizada"
// @Success 200 {object} entity.Resource "Recurso actualizado con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /resources [put]
func (ctrl *ResourceController) UpdateResource(c *gin.Context) {
	var resource entity.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.resourceService.UpdateResource(&resource); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resource)
}

// DeleteResource elimina un recurso por su ID
// @Summary Eliminar un recurso
// @Security      BearerAuth
// @Description Este endpoint permite eliminar un recurso específico usando su ID.
// @Tags recursos
// @Accept json
// @Produce json
// @Param id path uint true "ID del recurso"
// @Success 204 "Recurso eliminado con éxito"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /resources/{id} [delete]
func (ctrl *ResourceController) DeleteResource(c *gin.Context) {
	id, _ := util.StringToUint(c.Param("id"))
	if err := ctrl.resourceService.DeleteResource(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
