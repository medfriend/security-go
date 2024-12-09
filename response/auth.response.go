package response

import "security-go/entity"

type AuthResponse struct {
	Menus     []MenuResponse `json:"menus"`
	User      entity.User    `json:"user"`
	EntidadId uint           `json:"entidadId"`
}
