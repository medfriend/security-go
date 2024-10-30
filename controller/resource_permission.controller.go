package controller

import (
	"net/http"
	"security-go/entity"
	"security-go/service"
	"security-go/util"

	"github.com/gin-gonic/gin"
)

type ResourcePermissionController struct {
	ResourcePermissionService service.ResourcePermissionService
}

func NewResourcePermissionController(resourcePermissionService service.ResourcePermissionService) *ResourcePermissionController {
	return &ResourcePermissionController{
		ResourcePermissionService: resourcePermissionService,
	}
}

// CreateResourcePermission crea un nuevo ResourcePermission
// @Summary Crear un ResourcePermission
// @Description Este endpoint permite crear un nuevo ResourcePermission en el sistema.
// @Tags Recursos-permisos
// @Accept json
// @Produce json
// @Param resource body entity.ResourcePermission true "Información del // @Param resource body entity.ResourcePermission true "Información del ResourcePermission"
// @Success 201 {object} entity.ResourcePermission "ResourcePermission creado con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /medfri-security/resource_permission [post]"

func (ctrl *ResourcePermissionController) CreateResourcePermission(c *gin.Context) {
	var resourcePermission entity.ResourcePermission
	if err := c.ShouldBindJSON(&resourcePermission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.ResourcePermissionService.CreateResourcePermission(&resourcePermission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resourcePermission)
}

func (ctrl *ResourcePermissionController) GetResourcePermissionById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))

	resourcePermission, err := ctrl.ResourcePermissionService.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resourcePermission)
}

func (ctrl *ResourcePermissionController) UpdateResourcePermission(c *gin.Context) {
	var resourcePermission entity.ResourcePermission
	if err := c.ShouldBindJSON(&resourcePermission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.ResourcePermissionService.UpdateResourcePermission(&resourcePermission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resourcePermission)
}

func (ctrl *ResourcePermissionController) DeleteResourcePermission(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.ResourcePermissionService.DeleteResourcePermission(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}