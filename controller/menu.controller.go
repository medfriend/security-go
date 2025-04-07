package controller

import (
	"github.com/gin-gonic/gin"
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
// @Router /menu/createMenu [post]
func (ctrl *MenuController) CreateMenu(c *gin.Context) {
	var menu entity.Menu

	util.HandlerBindJson(c, &menu)
	util.HandlerInternalError(c, ctrl.MenuService.CreateMenu(&menu))
	util.HandlerCreatedSuccess(c, menu, menu.MenuID)
}

// GetChildByParentId
// @Summary      Obtener menus hijos desde el padre
// @Security      BearerAuth
// @Description  obtiene los hijos del menu padre por medio del id del padre
// @Tags         menus
// @Accept       json
// @Produce      json
// @Param        id  path      uint  true  "ID del menu"
// @Success      200 {object}  entity.Menu   "menu encontrado"
// @Router       /menu/childs-parent/{id} [get]
func (ctrl *MenuController) GetChildByParentId(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))

	menus, err := ctrl.MenuService.GetChildFromParentId(id)
	util.HandlerFoundSuccess(c, err, "menus hijos")
	util.HandlerCreatedSuccess(c, menus, 0)
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
// @Router       /menu/getMenuById/{id} [get]
func (ctrl *MenuController) GetMenuById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))

	menu, err := ctrl.MenuService.FindById(id)

	util.HandlerFoundSuccess(c, err, "menu")
	util.HandlerCreatedSuccess(c, menu, menu.MenuID)
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
// @Router /menu/updateMenu [put]
func (ctrl *MenuController) UpdateMenu(c *gin.Context) {
	var menu entity.Menu
	util.HandlerBindJson(c, &menu)
	util.HandlerInternalError(c, ctrl.MenuService.UpdateMenu(&menu))
	util.HandlerCreatedSuccess(c, menu, menu.MenuID)
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
	util.HandlerFoundSuccess(c, err, "menus padres")
	util.HandlerCreatedSuccess(c, menus, 0)
}

// FilterMenu obtener los menus por medio de la query
// @Summary filtro de menu
// @Security      BearerAuth
// @Description obtener menus por medio de la query
// @Param query path string true "campo de concidencia"
// @Tags menus
// @Accept json
// @Produce json
// @Success 200 {object} []entity.Menu "listado de menus"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /menu/filter/{query} [get]
func (ctrl *MenuController) FilterMenu(c *gin.Context) {
	query := c.Param("query")
	menus, err := ctrl.MenuService.FilterMenu(query)
	util.HandlerFoundSuccess(c, err, "menus concidentes")
	util.HandlerCreatedSuccess(c, menus, 0)
}

// GetChildsMenuByEntity
// @Summary menus con padres
// @Security      BearerAuth
// @Description obtener los menus con padres por en la entidad
// @Param entidadId path uint true "ID del menu"
// @Tags menus
// @Accept json
// @Produce json
// @Success 200 {object} []entity.Menu "listado de menus con padres"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /menu/childs/{entidadId} [get]
func (ctrl *MenuController) GetChildsMenuByEntity(c *gin.Context) {
	entidadId, _ := util.StringToUint(c.Param("entidadId"))

	menus, err := ctrl.MenuService.GetChildsMenuByEntity(entidadId)
	util.HandlerFoundSuccess(c, err, "menus hijos")
	util.HandlerCreatedSuccess(c, menus, 0)
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
// @Router /menu/deleteMenu/{id} [delete]
func (ctrl *MenuController) DeleteMenu(c *gin.Context) {
	id, _ := util.StringToUint(c.Param("id"))
	util.HandlerInternalError(c, ctrl.MenuService.DeleteMenu(id))
	util.HandlerNotContent(c, nil)
}
