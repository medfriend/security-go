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

func (ctrl *EntityController) GetEntityById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))
	Entity, err := ctrl.EntityService.GetEntityById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
		return
	}
	c.JSON(http.StatusOK, Entity)
}

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

func (ctrl *EntityController) DeleteEntity(c *gin.Context) {
	id, _ := util.StringToUint(c.Param("id"))
	if err := ctrl.EntityService.DeleteEntity(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
