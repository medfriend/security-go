package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"security-go/entity"
	"security-go/service"
)

type MenuController struct {
	MenuService service.MenuService
}

func NewMenuController(menuService service.MenuService) *MenuController {
	return &MenuController{
		MenuService: menuService,
	}
}

// CreateMenu crea un nuevo menu
// @Summary Crear un menu
// @Description Este endpoint permite crear un nuevo menu en el sistema.
// @Tags menus
// @Accept json
// @Produce json
// @Param resource body entity.Menu true "Información del menu"
// @Success 201 {object} entity.Menu "Menu creado con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /medfri-security/menu [post]
func (ctrl *MenuController) CreateMenu(c *gin.Context) {
	var menu entity.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := ctrl.MenuService.CreateMenu(&menu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, menu)
}
