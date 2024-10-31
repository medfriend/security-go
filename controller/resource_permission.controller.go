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

// CreateResourcePermission Create a new resource permission
// @Summary Create a new resource permission
// @Description This endpoint allows you to create a new resource permission in the system.
// @Tags recursos-permisos
// @Accept json
// @Produce json
// @Param resource body entity.ResourcePermission true "Resource permission information"
// @Success 201 {object} entity.ResourcePermission "Resource permission created successfully"
// @Failure 400 {object} map[string]string "Error in the request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /resource-permission [post]
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

// GetResourcePermissionById Get a resource permission by its ID
// @Summary Get a resource permission by ID
// @Description This endpoint returns the information of a specific resource permission given its ID.
// @Tags recursos-permisos
// @Accept json
// @Produce json
// @Param id path uint true "Resource permission ID"
// @Success 200 {object} entity.ResourcePermission "Resource permission found"
func (ctrl *ResourcePermissionController) GetResourcePermissionById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))

	resourcePermission, err := ctrl.ResourcePermissionService.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resourcePermission)
}

// UpdateResourcePermission Update a resource permission
// @Summary Update a resource permission
// @Description This endpoint allows you to update a resource permission in the system.
// @Tags recursos-permisos
// @Accept json
// @Produce json
// @Param resource body entity.ResourcePermission true "Updated resource permission information"
// @Success 200 {object} entity.ResourcePermission "Resource permission updated successfully"
// @Failure 400 {object} map[string]string "Error in the request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /resource-permission [put]
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

// DeleteResourcePermission Delete a resource permission
// @Summary Delete a resource permission
// @Description This endpoint allows you to delete an existing resource permission in the system.
// @Tags recursos-permisos
// @Accept json
// @Produce json
// @Param id path uint true "Resource permission ID"
// @Success 204 "Resource permission deleted successfully"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /resource-permission/{id} [delete]
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