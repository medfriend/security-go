package dto

type UserDTO struct {
	Usuario         string `json:"usuario"`
	Nombre1         string `json:"nombre_1"`
	Nombre2         string `json:"nombre_2"`
	ApellidoPaterno string `json:"apellido_paterno"`
	ApellidoMaterno string `json:"apellido_materno"`
	Clave           string `json:"clave"`
	Email           string `json:"email"`
	Edad            string `json:"edad"`
	Estado          bool   `json:"estado"`
}
