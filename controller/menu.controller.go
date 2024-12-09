package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"security-go/entity"
	"security-go/service"
	"security-go/util"
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
// @Security      BearerAuth
// @Description Este endpoint permite crear un nuevo menu en el sistema.
// @Tags menus
// @Accept json
// @Produce json
// @Param resource body entity.Menu true "Información del menu"
// @Success 201 {object} entity.Menu "Menu creado con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /menu [post]
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

// GetMenuById   obtiene un menu por su ID
// @Summary      Obtener un menu por ID
// @Security      BearerAuth
// @Description  Este endpoint devuelve la información de un menu específico dado su ID.
// @Tags         menus
// @Accept       json
// @Produce      json
// @Param        id  path      uint  true  "ID del menu"
// @Success      200 {object}  entity.Menu   "menu encontrado"
// @Router       /menu/{id} [get]
func (ctrl *MenuController) GetMenuById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))

	menu, err := ctrl.MenuService.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}
	c.JSON(http.StatusOK, menu)
}

// UpdateMenu actualiza un menu existente
// @Summary Actualizar un menu
// @Security      BearerAuth
// @Description Este endpoint permite actualizar la información de un menu existente.
// @Tags menus
// @Accept json
// @Produce json
// @Param resource body entity.Menu true "Información del menu actualizada"
// @Success 200 {object} entity.Menu "menu actualizado con éxito"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /menu [put]
func (ctrl *MenuController) UpdateMenu(c *gin.Context) {
	var menu entity.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.MenuService.UpdateMenu(&menu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, menu)
}

// GetParentsMenuByEntity obtener los menus padres de una entidad
// @Summary menus padres
// @Security      BearerAuth
// @Description obtener los menus padres de una entidad
// @Param entidadId path uint true "ID del menu"
// @Tags menus
// @Accept json
// @Produce json
// @Success 200 {object} []entity.Menu "listado de menus padres"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /menu/parents/{entidadId} [get]
func (ctrl *MenuController) GetParentsMenuByEntity(c *gin.Context) {
	entidadId, _ := util.StringToUint(c.Param("entidadId"))
	menus, err := ctrl.MenuService.GetParentsMenuByEntity(entidadId)
	//fmt.Println(menus, err)
	util.HandlerFoundSuccess(c, err, "menus padres")
	util.HandlerCreatedSuccess(c, menus)
}

// DeleteMenu elimina un menu por su ID
// @Summary Eliminar un menu
// @Security      BearerAuth
// @Description Este endpoint permite eliminar un menu específico usando su ID.
// @Tags menus
// @Accept json
// @Produce json
// @Param id path uint true "ID del menu"
// @Success 204 "menu eliminado con éxito"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /menu/{id} [delete]
func (ctrl *MenuController) DeleteMenu(c *gin.Context) {
	id, _ := util.StringToUint(c.Param("id"))
	if err := ctrl.MenuService.DeleteMenu(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
