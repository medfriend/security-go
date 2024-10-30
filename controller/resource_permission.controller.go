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