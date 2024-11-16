package mapper

import (
	"fmt"
	"security-go/entity"
	"security-go/response"
)

func MenuToMenuResponse(menu entity.Menu, hasSubmenu bool, permissions map[uint][]string) *response.MenuResponse {

	fmt.Println("entre al mapper")

	// Crear un `MenuResponse` básico
	menuResponse := &response.MenuResponse{
		MenuID:      menu.MenuID,
		Nombre:      menu.Nombre,
		Icono:       menu.Icono,
		Descripcion: menu.Descripcion,
		Submenus:    []*response.MenuResponse{},
		MenuPadreId: menu.MenuPadreID,
	}

	fmt.Println("paso 1")

	// Solo asignar `Recurso` si está disponible
	if menu.Recurso != nil {
		menuResponse.Recurso = menu.Recurso.Acceso
		menuResponse.RecursoID = menu.Recurso.ResourceID

		if permissions[menu.Recurso.ResourceID] != nil {
			//menuResponse.Permissions = permissions[menu.Recurso.ResourceID]
		}

	}

	fmt.Println("paso 2")

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

	fmt.Println("paso 3")

	return menuResponse
}
