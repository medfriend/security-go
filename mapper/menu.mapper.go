package mapper

import (
	"security-go/entity"
	"security-go/response"
)

func MenuToMenuResponse(menu entity.Menu, hasSubmenu bool, permissions map[uint][]string) *response.MenuResponse {

	// Crear un `MenuResponse` básico
	menuResponse := &response.MenuResponse{
		MenuID:      menu.MenuID,
		Nombre:      menu.Nombre,
		Icono:       menu.Icono,
		Descripcion: menu.Descripcion,
		Submenus:    []*response.MenuResponse{},
		MenuPadreId: menu.MenuPadreID,
	}

	// Solo asignar `Recurso` si está disponible
	if menu.Recurso != nil {
		menuResponse.Recurso = menu.Recurso.Acceso
		menuResponse.RecursoID = menu.Recurso.ResourceID

		if permissions[menu.Recurso.ResourceID] != nil {
			menuResponse.Permissions = permissions[menu.Recurso.ResourceID]
		}
	}

	if hasSubmenu && menu.MenuPadre != nil {
		// Crear una respuesta que representa al `MenuPadre`
		menuResponse = &response.MenuResponse{
			MenuID:      menu.MenuPadre.MenuID,
			Nombre:      menu.MenuPadre.Nombre,
			Icono:       menu.MenuPadre.Icono,
			Descripcion: menu.MenuPadre.Descripcion,
			Submenus:    []*response.MenuResponse{},
		}
	}

	return menuResponse
}
