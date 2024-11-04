package response

type MenuResponse struct {
	MenuID      uint            `json:"-"`
	Nombre      string          `json:"nombre"`
	Descripcion string          `json:"descripcion"`
	Submenus    []*MenuResponse `json:"submenus,omitempty"`
	Recurso     string          `json:"recurso,omitempty"`
	RecursoID   uint            `json:"-"`
	MenuPadreId *uint           `json:"-"`
	Permissions []string        `json:"permisos,omitempty"`
}
