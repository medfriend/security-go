package response

type MenuResponse struct {
	MenuID      uint            `json:"menu_id"`
	Nombre      string          `json:"nombre"`
	Descripcion string          `json:"descripcion"`
	Submenus    []*MenuResponse `json:"submenus,omitempty"`
	Recurso     string          `json:"recurso,omitempty"`
	MenuPadreId *uint           `json:"-"`
}
