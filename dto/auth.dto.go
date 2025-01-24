package dto

type AuthDTO struct {
	Usuario  uint   `json:"usuario"`
	Password string `json:"contraseña"`
}
