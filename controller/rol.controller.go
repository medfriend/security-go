package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"security-go/entity"
	"security-go/service"
	"security-go/util"
)

type RolController struct {
	RolService service.RolService
}

func NewRolController(RolService service.RolService) *RolController {
	return &RolController{
		RolService: RolService,
	}
}

// CreateRol crea un nuevo Rol
// @Summary Crear un Rol
// @Security      BearerAuth
// @Description Este endpoint permite crear un nuevo Rol en el sistema.
// @Tags   rols
// @Accept json
// @Produce json
// @Param resource body entity.Rol true "Información del Rol"
// @Success 201 {object} entity.Rol "Rol creado con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /rol [post]
func (ctrl *RolController) CreateRol(c *gin.Context) {
	var Rol entity.Rol
	if err := c.ShouldBindJSON(&Rol); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.RolService.CreateRol(&Rol); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Rol)
}

// GetRolById   obtiene un Rol por su ID
// @Summary      Obtener un Rol por ID
// @Security      BearerAuth
// @Description  Este endpoint devuelve la información de un Rol específico dado su ID.
// @Tags         rols
// @Accept       json
// @Produce      json
// @Param        id  path      uint  true  "ID del Rol"
// @Success      200 {object}  entity.Rol   "Rol encontrado"
// @Router       /rol/{id} [get]
func (ctrl *RolController) GetRolById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))

	Rol, err := ctrl.RolService.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rol not found"})
		return
	}
	c.JSON(http.StatusOK, Rol)
}

// UpdateRol actualiza un Rol existente
// @Summary Actualizar un Rol
// @Security      BearerAuth
// @Description Este endpoint permite actualizar la información de un Rol existente.
// @Tags rols
// @Accept json
// @Produce json
// @Param resource body entity.Rol true "Información del Rol actualizada"
// @Success 200 {object} entity.Rol "Rol actualizado con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /rol [put]
func (ctrl *RolController) UpdateRol(c *gin.Context) {
	var Rol entity.Rol
	if err := c.ShouldBindJSON(&Rol); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.RolService.UpdateRol(&Rol); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Rol)
}

// DeleteRol elimina un Rol por su ID
// @Summary Eliminar un Rol
// @Security      BearerAuth
// @Description Este endpoint permite eliminar un Rol específico usando su ID.
// @Tags rols
// @Accept json
// @Produce json
// @Param id path uint true "ID del Rol"
// @Success 204 "Rol eliminado con éxito"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /rol/{id} [delete]
func (ctrl *RolController) DeleteRol(c *gin.Context) {
	id, _ := util.StringToUint(c.Param("id"))
	if err := ctrl.RolService.DeleteRol(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
