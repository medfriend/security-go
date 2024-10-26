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

func (ctrl *PermisoController) CreatePermiso(c *gin.Context) {
	var permiso entity.Permiso
	if err := c.ShouldBindJSON(&permiso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.permisoService.CreatePermiso(&permiso); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, permiso)
}

func (ctrl *PermisoController) GetPermisoById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))
	permiso, err := ctrl.permisoService.GetPermisoById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Permiso not found"})
		return
	}
	c.JSON(http.StatusOK, permiso)
}

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

func (ctrl *PermisoController) DeletePermiso(c *gin.Context) {
	id, _ := util.StringToUint(c.Param("id"))
	if err := ctrl.permisoService.DeletePermiso(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
