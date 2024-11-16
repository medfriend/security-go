package response

type MenuResponse struct {
	MenuID      uint            `json:"-"`
	Nombre      string          `json:"label"`
	Descripcion string          `json:"descripcion"`
	Submenus    []*MenuResponse `json:"submenus,omitempty"`
	Recurso     string          `json:"route,omitempty"`
	RecursoID   uint            `json:"-"`
	MenuPadreId *uint           `json:"-"`
	Permissions []string        `json:"permisos,omitempty"`
	Icono       string          `json:"icon,omitempty"`
}
