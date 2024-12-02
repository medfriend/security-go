package mapper

import (
	"security-go/dto"
	"security-go/entity"
	"strconv"
	"time"
)

func UserDTOToUser(userDTO dto.UserDTO) (*entity.User, error) {
	usuarioID, err := strconv.ParseUint(userDTO.Usuario, 10, 32)
	edad, err := strconv.ParseUint(userDTO.Edad, 10, 32)

	if err != nil {
		return nil, err
	}

	return &entity.User{
		Usuario:         uint(usuarioID),
		Nombre1:         userDTO.Nombre1,
		Nombre2:         userDTO.Nombre2,
		ApellidoPaterno: userDTO.ApellidoPaterno,
		ApellidoMaterno: userDTO.ApellidoMaterno,
		Clave:           userDTO.Clave,
		Email:           userDTO.Email,
		FechaCreacion:   time.Now(),
		Edad:            uint(edad),
		Activo:          userDTO.Estado,
	}, nil
}
