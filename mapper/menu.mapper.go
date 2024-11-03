package mapper

import (
	"security-go/entity"
	"security-go/response"
)

func MenuToMenuResponse(menu entity.Menu, hasSubmenu bool) *response.MenuResponse {

	menuResponse := &response.MenuResponse{}

	if !hasSubmenu {
		menuResponse = &response.MenuResponse{
			MenuID:      menu.MenuID,
			Nombre:      menu.Nombre,
			Descripcion: menu.Descripcion,
			Recurso:     menu.Recurso.Acceso,
		}
	}

	if hasSubmenu {
		menuResponse = &response.MenuResponse{
			MenuID:      menu.MenuPadre.MenuID,
			Nombre:      menu.MenuPadre.Nombre,
			Descripcion: menu.MenuPadre.Descripcion,
			Submenus:    []*response.MenuResponse{},
		}
	}

	return menuResponse
}
