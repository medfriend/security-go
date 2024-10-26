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

func (ctrl *ResourceController) GetResourceById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))
	resource, err := ctrl.resourceService.GetResourceById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		return
	}
	c.JSON(http.StatusOK, resource)
}

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

func (ctrl *ResourceController) DeleteResource(c *gin.Context) {
	id, _ := util.StringToUint(c.Param("id"))
	if err := ctrl.resourceService.DeleteResource(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}